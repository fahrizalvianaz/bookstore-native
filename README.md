# Simple Go Web API

A lightweight web API built with native Go packages implementing a Go standart project pattern with repository, service, and handler layers.

## Features

- RESTful API endpoints
- Clean architecture implementation
- PostgreSQL database integration
- User authentication (register/login)
- Book management (CRUD operations)

## Project Structure

```
simple-project-native/
├── cmd/
│   └── myapp1/
│       ├── config/        # Configuration handling
│       ├── handler/       # HTTP handlers
│       ├── model/         # Data models
│       │   └── dto/       # Data Transfer Objects
│       ├── repository/    # Database operations
│       ├── service/       # Business logic
│       └── main.go        # Application entry point
├── pkg/
│   └── utils/             # Utility functions
└── README.md
```

## Prerequisites

- Go 1.22+ (for 2023 HTTP handler syntax)
- PostgreSQL database

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/simple-project-native.git
   cd simple-project-native
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure database connection in your environment or config file

4. Run the application:
   ```bash
   go run cmd/myapp1/main.go
   ```

## API Endpoints

### Books

- **Create a book**
  - `POST /book`
  - Request body: `{ "title": "Book Title", "author": "Author Name" }`

- **Get a book by ID**
  - `GET /book/{id}`

- **Get all books**
  - `GET /books`

### Users

- **Register a new user**
  - `POST /register`
  - Request body: `{ "nama": "Full Name", "email": "user@example.com", "username": "username", "password": "password" }`

- **User login**
  - `POST /login`
  - Request body: `{ "username": "username", "password": "password" }`
  - Returns a success message if authentication is successful

## Architecture

This project follows Go standart project pattern:

1. **Handler Layer**: Manages HTTP requests and responses
2. **Service Layer**: Contains business logic
3. **Repository Layer**: Handles data access and persistence

## Database Schema

### Books Table
```sql
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL
)
```

### Users Table
```sql
CREATE TABLE IF NOT EXISTS customer (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
)
```

## Security Features

- Password hashing using bcrypt
- Input validation
- Proper error handling

