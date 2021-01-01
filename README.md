# go-to-cafe

This is a sample repository of modular monolith architecture.


The main features are below.
- High Module Independency
    - Each domain-specific module is independent.
    - There are strong constraints for crossing domain boundaries.
- High Deployability
    - Low deployment costs due to a single server.

# Architecture
```
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   └── api
│       └── server # This is a main entry point.
├── db
│   ├── migrations
│   └── schema.sql
├── docker-compose.yaml
├── domain # This is the directory with the independent domain modules.
│   ├── auth
│   │   ├── adapter
│   │   ├── pkg
│   │   └── usecase
│   ├── cafe
│   │   ├── adapter
│   │   ├── pkg
│   │   └── usecase
│   ├── search
│   │   ├── adapter
│   │   ├── pkg
│   │   └── usecase
│   └── user
│       ├── adapter
│       ├── pkg
│       └── usecase
├── go.mod
├── go.sum
├── pkg
│   ├── apperror
│   ├── boundary # These are the interfaces to use when crossing the domain boundary. (by using `singleton`, `observer`, `go interface`)
│   │   ├── create_cafe
│   │   │   ├── event.go
│   │   │   └── notifer.go
│   │   └── sign_up
│   │       ├── event.go
│   │       └── notifer.go
│   ├── config
│   ├── elasticsearch
│   ├── mysql
│   ├── presenter
│   └── registry
└── server
```

# Usage

```
$ make migrate
$ make run
```
