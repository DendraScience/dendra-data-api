package types

import (
	"fmt"
)

type PathError struct {
	Path string
}

func (e *PathError) Error() string {
	return fmt.Sprintf("path not valid %q", e.Path)
}

type ParamsError struct {
	Detail string
}

func (e *ParamsError) Error() string {
	return fmt.Sprintf("params not valid: %s", e.Detail)
}
