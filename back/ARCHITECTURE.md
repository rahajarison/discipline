# Architecture Documentation

## Project Structure

```
back/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── handlers/               # HTTP handlers (presentation layer)
│   │   ├── user_handler.go
│   │   ├── match_handler.go
│   │   ├── round_handler.go
│   │   └── action_handler.go
│   └── services/               # Business logic layer
│       └── user_service.go
├── pkg/
│   ├── database/               # Database connection and configuration
│   │   └── database.go
│   └── validation/             # Reusable validation functions
│       └── validation.go
├── api/
│   ├── router.go               # Route definitions
│   └── middleware.go           # HTTP middleware
├── models/                     # Data models
│   ├── user.go
│   ├── match.go
│   ├── round.go
│   ├── action.go
│   └── base.go
└── API_ENDPOINTS.md            # API documentation
```

## Architecture Layers

### 1. Presentation Layer (`internal/handlers/`)
- **Purpose**: Handle HTTP requests and responses
- **Responsibilities**:
  - Parse request data
  - Validate input
  - Call business logic services
  - Format responses
  - Handle errors

### 2. Business Logic Layer (`internal/services/`)
- **Purpose**: Implement business rules and logic
- **Responsibilities**:
  - Business operations
  - Data processing
  - Business validation
  - Orchestrate data access

### 3. Data Access Layer (`pkg/database/`)
- **Purpose**: Handle database operations
- **Responsibilities**:
  - Database connection
  - Query execution
  - Data mapping

### 4. Models (`models/`)
- **Purpose**: Define data structures
- **Responsibilities**:
  - Entity definitions
  - Database schema mapping
  - Data validation tags

## Design Patterns

### 1. Separation of Concerns
Each layer has a specific responsibility:
- **Handlers**: HTTP concerns
- **Services**: Business logic
- **Database**: Data persistence

### 2. Dependency Injection
Handlers receive their dependencies (database, services) through constructors:
```go
func NewUserHandler(db *gorm.DB) *UserHandler {
    return &UserHandler{db: db}
}
```

### 3. Single Responsibility Principle
Each handler file is responsible for one entity:
- `user_handler.go` → User operations only
- `match_handler.go` → Match operations only
- etc.

## Benefits of This Architecture

### 1. Maintainability
- **Easy to locate code**: Each entity has its own file
- **Clear separation**: Business logic separate from HTTP handling
- **Modular design**: Changes to one entity don't affect others

### 2. Testability
- **Isolated testing**: Each layer can be tested independently
- **Mock dependencies**: Easy to mock database and services
- **Unit tests**: Business logic can be tested without HTTP

### 3. Scalability
- **Add new entities**: Just create new handler and service files
- **Modify existing**: Changes are isolated to specific files
- **Team development**: Different developers can work on different entities

### 4. Reusability
- **Services**: Can be used by different handlers
- **Validation**: Reusable across the application
- **Database**: Centralized database operations

## File Organization Principles

### 1. `internal/` vs `pkg/`
- **`internal/`**: Code specific to this application
- **`pkg/`**: Code that could be reused in other projects

### 2. Handler Organization
- One file per entity
- All CRUD operations in the same file
- Related operations grouped together

### 3. Service Organization
- Business logic separated from HTTP concerns
- Reusable across different handlers
- Easy to test independently

## Future Enhancements

### 1. Repository Pattern
Add repository layer between services and database:
```
services → repositories → database
```

### 2. DTOs (Data Transfer Objects)
Add DTOs for request/response formatting:
```
HTTP → DTOs → Models → Database
```

### 3. Middleware Chain
Add more middleware for:
- Authentication
- Authorization
- Rate limiting
- Request logging

### 4. Error Handling
Centralized error handling with custom error types and proper HTTP status codes.

## Code Examples

### Handler Structure
```go
type UserHandler struct {
    db *gorm.DB
}

func (h *UserHandler) CreateUser(c echo.Context) error {
    // 1. Parse request
    // 2. Validate input
    // 3. Call business logic
    // 4. Return response
}
```

### Service Structure
```go
type UserService struct {
    db *gorm.DB
}

func (s *UserService) CreateUser(user *models.User) error {
    // Business logic here
    return s.db.Create(user).Error
}
```

### Router Setup
```go
func SetupRoutes(e *echo.Echo, userHandler *handlers.UserHandler, ...) {
    users := v1.Group("/users")
    users.POST("", userHandler.CreateUser)
    // ...
}
```
