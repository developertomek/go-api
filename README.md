# GoFEM API

A fitness and exercise management REST API built with Go. Track workouts, manage users, and handle authentication.

## Features

- User authentication and management
- Workout tracking
- Protected endpoints with token authentication
- PostgreSQL database integration

## Prerequisites

- Go 1.20 or higher
- PostgreSQL
- [goose](https://github.com/pressly/goose) for database migrations

## Quick Start

1. Clone the repository
2. Setup PostgreSQL database
3. Run migrations:
```bash
goose postgres "postgres://user:password@localhost:5432/dbname?sslmode=disable" up
```
4. Run the server:
```bash
go run main.go
```

## API Endpoints

### Authentication
- `POST /users` - Register new user
- `POST /tokens/authentication` - Login and get auth token

### Workouts (Protected Routes)
- `GET /workouts/{id}` - Get workout by ID
- `POST /workouts` - Create new workout
- `PUT /workouts/{id}` - Update workout
- `DELETE /workouts/{id}` - Delete workout

### Health Check
- `GET /health` - API health check

## Example Usage

### Register a new user
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "secret123"
  }'
```

### Create a new workout (with auth token)
```bash
curl -X POST http://localhost:8080/workouts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_AUTH_TOKEN" \
  -d '{
    "title": "Morning Run",
    "description": "5km run in the park",
    "duration_minutes": 30,
    "calories_burned": 300
  }'
```

## Database Schema

The API uses two main tables:
- `users` - Stores user information and credentials
- `workouts` - Stores workout details linked to users

## Error Handling

The API returns consistent error responses in the following format:
```json
{
    "error": "error message here"
}
```

## Security

- Passwords are hashed before storage
- Tokens for authentication
- Protected routes require valid authentication
- User-specific workout access control

## Development

The project follows a clean architecture pattern with the following structure:
- `/internal/api` - HTTP handlers
- `/internal/store` - Database interactions
- `/internal/middleware` - HTTP middleware
- `/internal/routes` - Route definitions
- `/migrations` - Database migrations
