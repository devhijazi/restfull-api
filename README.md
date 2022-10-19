# restfull-api

Restful API made in goLang using principles of Clean Code and Clean Architecture

# Requisitos

- Clean Code
- Clean Arch
- Swagger
- Docker
- Dependency Injection
- Postgres (gorm)
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
│   ├── docker-compose.yml
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
│   ├── mocks
└── README.md
```

# Overview

- Cadastro de Usuários
  - Autenticação (jwt)
  - CRUD
- Cadastro de Loja (Relacionado ao usuário)
  - Autenticado
  - CRUD de Loja
- Cadastro de Produtos
  - Autenticado
  - CRUD de Produtos
  - Paginação
- Entrada/Saida
- Balance
