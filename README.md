# Simple golang with gin

## Getting started

```bash
# 1. Clone the repository or click on "Use this template" button.
git clone https://github.com/devlorz/pea-cpm.git

# 2. Enter your newly-cloned folder.
cd pea-cpm

# 3. Run
go run main.go
```

# Setup and development

- [Setup and development](#setup-and-development)
  - [First-time setup](#first-time-setup)
  - [Installation](#installation)
    - [Configuration](#configuration)
  - [Generators](#generators)
  - [Web Framework](#web-framework)
  - [Dependency](#dependency)

## First-time setup

Make sure you have the following installed:

- [Golang](https://go.dev/dl/)

### Configuration

Before start fill correct configurations in `.env` file.

### Simple environment

```env
#########################
# APP DEVELOPMENT ENV
#########################
GO_PORT = 3000
GO_ENV = test

#########################
# PG DEVELOPMENT ENV
#########################
DATABASE_URI_PROD = postgres://admin:admin@123@localhost:5432/postgres
DATABASE_URI_DEV = postgres://admin:admin@123@localhost:5432/postgres

#########################
# JWT DEVELOPMENT ENV
########################
JWT_SECRET = 2c46f7651a176169c8d2f7aee60cce3da874501d

```

## Web Framework

- gin-gonic (https://github.com/gin-gonic/gin)

## Dependency

- godotenv (https://github.com/joho/godotenv)
- jwt-go (https://github.com/dgrijalva/jwt-go)
- validator (https://github.com/go-playground/validator)
- gorm (https://gorm.io/index.html)
- go-playground-converter (https://github.com/restuwahyu13/go-playground-converter)
