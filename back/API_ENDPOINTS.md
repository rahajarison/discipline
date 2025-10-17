# API Endpoints Documentation

## Base URL
```
http://localhost:8080/api/v1
```

## Health Check
- **GET** `/health` - Check API status

## Users

### CRUD Operations
- **POST** `/users` - Create a new user
- **GET** `/users` - Get all users (with pagination)
- **GET** `/users/:id` - Get a user by ID
- **PUT** `/users/:id` - Update a user
- **DELETE** `/users/:id` - Delete a user

### Relations
- **GET** `/users/:userId/matches` - Get all matches for a user

### Query Parameters for `/users`
- `page` (int) - Page number (default: 1)
- `limit` (int) - Number of items per page (default: 10)

## Matches

### CRUD Operations
- **POST** `/matches` - Create a new match
- **GET** `/matches` - Get all matches (with pagination)
- **GET** `/matches/:id` - Get a match by ID
- **PUT** `/matches/:id` - Update a match
- **DELETE** `/matches/:id` - Delete a match

### Query Parameters for `/matches`
- `page` (int) - Page number (default: 1)
- `limit` (int) - Number of items per page (default: 10)

## Rounds

### CRUD Operations (nested under matches)
- **POST** `/matches/:matchId/rounds` - Create a new round
- **GET** `/matches/:matchId/rounds` - Get all rounds for a match
- **GET** `/matches/:matchId/rounds/:id` - Get a round by ID
- **PUT** `/matches/:matchId/rounds/:id` - Update a round
- **DELETE** `/matches/:matchId/rounds/:id` - Delete a round

### Direct Access (alternative)
- **GET** `/rounds/:id` - Get a round by ID
- **PUT** `/rounds/:id` - Update a round
- **DELETE** `/rounds/:id` - Delete a round

## Actions

### CRUD Operations (nested under rounds)
- **POST** `/matches/:matchId/rounds/:roundId/actions` - Create a new action
- **GET** `/matches/:matchId/rounds/:roundId/actions` - Get all actions for a round
- **GET** `/matches/:matchId/rounds/:roundId/actions/:id` - Get an action by ID
- **PUT** `/matches/:matchId/rounds/:roundId/actions/:id` - Update an action
- **DELETE** `/matches/:matchId/rounds/:roundId/actions/:id` - Delete an action

### Relations
- **GET** `/matches/:matchId/rounds/:roundId/actions/player/:player` - Get actions for a player (P1 or P2)

### Direct Access (alternative)
- **GET** `/actions/:id` - Get an action by ID
- **PUT** `/actions/:id` - Update an action
- **DELETE** `/actions/:id` - Delete an action

## Usage Examples

### Create a user
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "player1",
    "email": "player1@example.com",
    "role": "player",
    "passwordHash": "hashedpassword"
  }'
```

### Create a match
```bash
curl -X POST http://localhost:8080/api/v1/matches \
  -H "Content-Type: application/json" \
  -d '{
    "player1Id": "uuid-player1",
    "player2Id": "uuid-player2",
    "date": "2024-01-01T10:00:00Z",
    "replayUrl": "https://example.com/replay",
    "notes": "Match notes"
  }'
```

### Create a round
```bash
curl -X POST http://localhost:8080/api/v1/matches/:matchId/rounds \
  -H "Content-Type: application/json" \
  -d '{
    "number": 1
  }'
```

### Create an action
```bash
curl -X POST http://localhost:8080/api/v1/matches/:matchId/rounds/:roundId/actions \
  -H "Content-Type: application/json" \
  -d '{
    "type": "hit",
    "category": "attack",
    "hitContext": "high",
    "player": "P1",
    "timestamp": "2024-01-01T10:05:00Z"
  }'
```

## Response Codes

- **200** - Success
- **201** - Created successfully
- **400** - Invalid request
- **404** - Resource not found
- **500** - Internal server error

## Error Response Structure

```json
{
  "error": "Descriptive error message"
}
```

## Pagination

Endpoints that return lists support pagination via query parameters `page` and `limit`.

Example:
```
GET /api/v1/users?page=2&limit=5
```

## Automatic Relations

Some endpoints automatically return relations:
- `/matches/:id` includes `Player1` and `Player2` data
- `/rounds/:id` includes `Match` data
- `/actions/:id` includes `Round` data
