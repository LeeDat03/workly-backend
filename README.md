## SETUP

setup project with env like env.example

-   Setup Docker container for pg + redis

```
docker compose up -d
```

-   Build and run app

```
air
```

## 🔄 Data Flow (Clean Architecture)

```
┌─────────────┐
│   Client    │
└──────┬──────┘
       │
       ▼
┌─────────────────────────────────────────┐
│         HTTP Handler Layer              │
│  ┌────────────────────────────────┐    │
│  │  Router → Middleware → Handler │    │
│  └────────────────────────────────┘    │
└──────────────────┬──────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────┐
│          UseCase Layer                  │
│  ┌────────────────────────────────┐    │
│  │  Business Logic & Validation   │    │
│  └────────────────────────────────┘    │
└──────────────────┬──────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────┐
│         Repository Layer                │
│  ┌────────────────────────────────┐    │
│  │    Data Access & Queries       │    │
│  └────────────────────────────────┘    │
└──────────────────┬──────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────┐
│       Infrastructure Layer              │
│  ┌────────────────────────────────┐    │
│  │  PostgreSQL │ Redis │ SMTP     │    │
│  └────────────────────────────────┘    │
└─────────────────────────────────────────┘
```

## 🏗️ Architecture Layers

### 1. **Handler Layer** (`internal/handler/`)

-   Receives HTTP requests
-   Validates input using DTOs
-   Calls appropriate UseCase
-   Returns formatted responses

### 2. **UseCase Layer** (`internal/usecase/`)

-   Contains business logic
-   Orchestrates data flow
-   Independent of external frameworks
-   Calls repositories for data operations

### 3. **Repository Layer** (`internal/repository/`)

-   Handles data persistence
-   Abstracts database operations
-   Implements data access interfaces

### 4. **Infrastructure Layer** (`internal/infrastructure/`)

-   External dependencies (Database, Cache, Email)
-   Third-party integrations
-   Configuration and setup

### 5. **Domain Layer** (`internal/domain/`)

-   Core business entities
-   Domain-specific errors
-   Business rules and validation

## 🚀 Request Flow Example

```
1. Client sends POST /api/v1/auth/register
                ↓
2. Router matches route → auth_handler.Register()
                ↓
3. Handler validates DTO → auth_dto.RegisterRequest
                ↓
4. Handler calls UseCase → auth_usecase.Register()
                ↓
5. UseCase processes business logic
                ↓
6. UseCase calls Repository → user_db.Create()
                ↓
7. Repository executes SQL query → PostgreSQL
                ↓
8. Response flows back up the chain
                ↓
9. Handler formats response → JSON
                ↓
10. Client receives response
```
