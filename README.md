# restfull-api

Postgres | Gin | GoLang

# Requisitos

- Clean Code
- Clean Arch
- Swagger
- Docker
- Dependency Injection
- Postgres
- GraphQL
- Errors (Tratativa de Errors)
- Unity Tests
- Charge Tests

# Architecture

```
├── cmd
│   ├── main.go
│   ├── server.go
│   ├── injection.go
│   └── route.go
├── internal
│   ├── controllers
│   │   └── file.go
│   ├── middlewares
│   │   └── file.go
│   ├── models
│   │   └── file.go
│   ├── repository
│   │   └── file.go
│   ├── services
│   │   └── file.go
│   └── routes
│       └── users.go
├── docker
│   └── Dockerfile
├── docs
│   └── [swagger genenrate]
├── migrations
│   └── file.go
├── pkg
│   ├── database
│   │   ├── connection.go
│   ├── helpers
│   ├── errors
│   ├── validations
│   ├── utils
└── README.md
```
