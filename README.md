# Dendra Services

Dendra's application services and API with in-flight data processing middleware. Realized as a collection of interconnected gRPC services with a REST API gateway.

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
- See Makefile for more.
