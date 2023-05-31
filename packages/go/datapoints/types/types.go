package types

import (
	"fmt"
	"sort"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/dendrascience/dendra-data-api/release/go/v3"
)

var (
	maxTime = time.Date(2200, time.January, 2, 0, 0, 0, 0, time.UTC)
	minTime = time.Date(1800, time.January, 2, 0, 0, 0, 0, time.UTC)
)

type Interval struct {
	Start     time.Time
	End       time.Time
	LeftOpen  bool
	RightOpen bool
}

func NewEmptyInterval() Interval {
	return Interval{
		Start:     minTime,
		End:       minTime,
		LeftOpen:  true,
		RightOpen: true,
	}
}

func NewQueryInterval() Interval {
	return Interval{
		Start:     minTime,
		End:       maxTime,
		LeftOpen:  false,
		RightOpen: true,
	}
}

func (self Interval) Intersect(other Interval) Interval {
	if self.Start.After(other.End) || self.End.Before(other.Start) {
		return NewEmptyInterval()
	}

	start := self.Start
	leftOpen := self.LeftOpen || other.LeftOpen

	if self.Start.Before(other.Start) {
		start = other.Start
		leftOpen = other.LeftOpen
	} else if self.Start.After(other.Start) {
		start = self.Start
		leftOpen = self.LeftOpen
	}

	end := self.End
	rightOpen := self.RightOpen || other.RightOpen

	if self.End.Before(other.End) {
		end = self.End
		rightOpen = self.RightOpen
	} else if self.End.After(other.End) {
		end = other.End
		rightOpen = other.RightOpen
	}

	return Interval{
		Start:     start,
		End:       end,
		LeftOpen:  leftOpen,
		RightOpen: rightOpen,
	}
}

func (self Interval) IsEmpty() bool {
	return self.End.Before(self.Start) ||
		(self.End.Equal(self.Start) && (self.LeftOpen || self.RightOpen))
}

func BucketTsFn(width string) (func(time.Time, time.Time) time.Time, error) {
	if width == "" {
		return func(ts time.Time, currTs time.Time) time.Time {
			if currTs.IsZero() {
				return ts
			}
			return currTs
		}, nil
	}

	// fixed calendar durations
	if width == "1y" {
		return func(ts time.Time, _ time.Time) time.Time {
			return time.Date(ts.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
		}, nil
	}
	if width == "1mo" {
		return func(ts time.Time, _ time.Time) time.Time {
			return time.Date(ts.Year(), ts.Month(), 1, 0, 0, 0, 0, time.UTC)
		}, nil
	}
	if width == "1d" {
		return func(ts time.Time, _ time.Time) time.Time {
			return time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.UTC)
		}, nil
	}

	// fallback to time durations
	duration, err := time.ParseDuration(width)
	if err != nil {
		return nil, fmt.Errorf("parse bucket width error: %w", err)
	}
	return func(ts time.Time, _ time.Time) time.Time {
		return ts.Truncate(duration)
	}, nil
}

func MergeConfig(datastream *pb.Datastream, reverse bool) []*pb.DatapointsConfigInstance {
	config := datastream.GetDatapointsConfig()
	if config == nil {
		return make([]*pb.DatapointsConfigInstance, 0)
	}

	builtConfig := datastream.GetDatapointsConfigBuilt()
	if builtConfig != nil && len(builtConfig) > 0 {
		config = builtConfig
	}

	refdConfig := datastream.GetDatapointsConfigRefd()

	newConfig := make([]*pb.DatapointsConfigInstance, 0, len(config))

	//
	// efficiently merge config instances in a linear traversal by evaluating
	// each instance's date/time interval [begins_at, ends_before)
	//

	for _, inst := range config {
		actions := inst.GetActions()
		if actions != nil {
			if actions.GetExclude() {
				continue
			}
		}

		newInst := proto.Clone(inst).(*pb.DatapointsConfigInstance)

		if newInst.Ref != nil {
			refIndex := int(newInst.GetRef())

			if refIndex < len(refdConfig) {
				refInst := refdConfig[refIndex]

				newInst.Connection = refInst.GetConnection()
				newInst.Path = refInst.GetPath()
				newInst.Params = proto.Clone(refInst.GetParams()).(*structpb.Struct)
			}
		}

		if newInst.GetConnection() != "" {
			// TODO: handle connection-backed queries or deprecate, for now ingore
			continue
		}
		if newInst.BeginsAt == nil {
			newInst.BeginsAt = timestamppb.New(minTime)
		}
		if newInst.EndsBefore == nil {
			newInst.EndsBefore = timestamppb.New(maxTime)
		}

		newConfig = append(newConfig, newInst)
	}

	// sort instances by beginsAt
	sort.SliceStable(newConfig, func(i, j int) bool {
		return newConfig[i].GetBeginsAt().AsTime().Before(newConfig[j].GetBeginsAt().AsTime())
	})

	stack := make([]*pb.DatapointsConfigInstance, 0, len(newConfig))

	for _, inst := range newConfig {
		instBeginsAt := inst.GetBeginsAt().AsTime()
		instEndsBefore := inst.GetEndsBefore().AsTime()

		if !instEndsBefore.After(instBeginsAt) {
			// exclude: inverted interval
			continue
		}

		if len(stack) == 0 {
			// init stack
			stack = append(stack, inst)
		} else {
			top := stack[len(stack)-1]
			topBeginsAt := top.GetBeginsAt().AsTime()
			topEndsBefore := top.GetEndsBefore().AsTime()

			if instBeginsAt.After(topEndsBefore) || instBeginsAt.Equal(topEndsBefore) {
				stack = append(stack, inst)
			} else if instEndsBefore.Before(topEndsBefore) || instEndsBefore.Equal(topEndsBefore) {
				// exclude: instance interval is within top interval
				continue
			} else if instBeginsAt.Equal(topBeginsAt) {
				stack[len(stack)-1] = inst
			} else {
				top.EndsBefore = timestamppb.New(instBeginsAt)
				stack = append(stack, inst)
			}
		}
	}

	if reverse {
		for i := 0; i < len(stack)/2; i++ {
			j := len(stack) - i - 1
			stack[i], stack[j] = stack[j], stack[i]
		}
	}

	return stack
}
