# Dendra Data API

Dendra's data query API with in-flight data processing middleware. Realized as a collection of interconnected gRPC services with a REST API facade.

This is a multi-language monorepo. 

## Project structure

```
.
├── packages   # Source for all services organized by language and purpose
├── proto      # The protobuf files for all messages and services
└── release    # Output from protoc for inclusion, organized by language
```

## Notes

- This is a proof of concept and work in progress.
- A REST API facade will be developed in Node to handle incoming (signed) data query calls. This endeavors to support the multi data-site concept, with a single meta-site.


## Service types

- DerivedDatastreamService — Higher-level data manipulation services to handle packaged calculations with vetted libraries, e.g. MetPy/Python. Calculations are registered on the service bus, and may utilize the AggregationService, DatastreamQueryService and ProviderQueryService.
- AggregationService — Performs core/basic aggregations (e.g. Count, Max, Min, Mean, Stddev, Sum, Wysum). Calls the ProviderQueryService to fetch datapoints or base aggregates (aggregates that the provider can quickly determine or can push down to the engine for processing).
- DatastreamQueryService — Performs multi-datastream queries and transformations. Calls the ProviderQueryService to assemble in-memory multi-datastream tables for stream-based transformations.
- ProviderQueryService — Calls different ProviderServices to fetch time-series data based on the configured source for a given datastream config instance. May call different transformers or converters to post-process and shape the result.
- ConverterService — Called perform units conversion using metadata. Different converters can be built in a variety of languages to support vetted conversion libraries, e.g. Pint/Python.
- ProviderService — Fetches time-series data from a data source, e.g. Influx or SQL. Currently only [Influx](packages/go/provider/influx).
