# Library Management API

A robust REST API for managing a fictional public library built with Go, PostgreSQL, Redis caching, and containerized with Docker.

## Architecture Overview

```
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── database/
│   │   ├── postgres.go          # PostgreSQL connection
│   │   └── migrations/
│   │       └── 001_create_books.sql
│   ├── cache/
│   │   └── redis.go             # Redis cache implementation
│   ├── models/
│   │   └── book.go              # Data models
│   ├── handlers/
│   │   └── book_handler.go      # HTTP handlers
│   ├── services/
│   │   └── book_service.go      # Business logic
│   └── middleware/
│       └── middleware.go        # HTTP middleware
├── pkg/
│   └── utils/
│       └── response.go          # Response utilities
├── docker-compose.yml           # Multi-container setup
├── Dockerfile                   # Go application container
├── go.mod                      # Go dependencies
├── go.sum
└── README.md                   # This file
```

## Features

- **Complete CRUD Operations**: Create, Read, Update, Delete books
- **High Performance**: Redis caching with 15-minute TTL
- **Database**: PostgreSQL with proper migrations
- **REST API**: Clean RESTful endpoints using Gin framework
- **Containerized**: Docker & Docker Compose setup
- **Production Ready**: Proper error handling, logging, and validation
- **Clean Architecture**: Separation of concerns with layers

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/books` | List all books |
| GET | `/api/v1/books/:id` | Get book by ID |
| POST | `/api/v1/books` | Create new book |
| PUT | `/api/v1/books/:id` | Update book |
| DELETE | `/api/v1/books/:id` | Delete book |
| GET | `/health` | Health check |

## Tech Stack

- **Language**: Go 1.21
- **Framework**: Gin HTTP Framework
- **Database**: PostgreSQL
- **Cache**: Redis
- **Containerization**: Docker & Docker Compose
- **ORM**: GORM
- **Migration**: golang-migrate

## Quick Start

### Prerequisites
- Docker and Docker Compose
- Go 1.21+ (for local development)

### Using Docker Compose (Recommended)

1. **Clone the repository**
```bash
git clone 
cd library-management-api
```

2. **Create environment file**
```bash
cp .env.example .env
# Edit .env with your database credentials
```

3. **Start the application**
```bash
docker-compose up --build
```

4. **Test the API**
```bash
curl http://localhost:8080/health
```

### Local Development

1. **Install dependencies**
```bash
go mod download
```

2. **Set environment variables**
```bash
export DATABASE_URL="postgresql://postgres.wywkucanulrrkqgexwcp:Tejas%402001@aws-0-ap-south-1.pooler.supabase.com:5432/postgres"
export REDIS_URL="redis://default:hDbjDRpv9yi892LytkwuAs1yrKSw8cjL@redis-14159.c206.ap-south-1-1.ec2.redns.redis-cloud.com:14159"
export ENVIRONMENT="development"
```

3. **Run the application**
```bash
go run cmd/server/main.go
```

## API Usage Examples

### Create a Book
```bash
curl -X POST http://localhost:8080/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Go Programming Language",
    "author": "Alan Donovan",
    "isbn": "978-0134190440",
    "published_year": 2015,
    "genre": "Programming",
    "available_copies": 5
  }'
```

### Get All Books
```bash
curl http://localhost:8080/api/v1/books
```

### Get Book by ID
```bash
curl http://localhost:8080/api/v1/books/1
```

### Update Book
```bash
curl -X PUT http://localhost:8080/api/v1/books/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Go Programming Language - Updated",
    "author": "Alan Donovan & Brian Kernighan",
    "isbn": "978-0134190440",
    "published_year": 2015,
    "genre": "Programming",
    "available_copies": 3
  }'
```

### Delete Book
```bash
curl -X DELETE http://localhost:8080/api/v1/books/1
```

## Architecture Decisions

### Caching Strategy
- **Redis TTL**: 15 minutes for individual books and book lists
- **Cache Keys**: Structured as `book:id` and `books:all`
- **Cache-Aside Pattern**: Check cache first, fallback to database
- **Cache Invalidation**: On create, update, delete operations

### Database Design
- **PostgreSQL**: Reliable, ACID-compliant relational database
- **GORM**: Type-safe ORM with migration support
- **Connection Pooling**: Optimized connection management

### API Design
- **RESTful**: Follows REST principles
- **JSON**: Standard request/response format
- **HTTP Status Codes**: Proper status code usage
- **Error Handling**: Consistent error response format

### Security Considerations
- **Environment Variables**: Sensitive data stored securely
- **Input Validation**: Request validation using Gin binding
- **Error Messages**: Safe error messages without data leakage

## Docker Configuration

### Application Container
- **Multi-stage build**: Optimized image size
- **Non-root user**: Security best practice
- **Health checks**: Container health monitoring

### Docker Compose
- **Service isolation**: Separate containers for app and database
- **Volume management**: Persistent data storage
- **Network configuration**: Service communication

## Testing

Run tests:
```bash
go test ./...
```

Run with coverage:
```bash
go test -cover ./...
```

## Performance Features

- **Connection Pooling**: Database connection optimization
- **Caching Layer**: Redis for fast data retrieval
- **JSON Streaming**: Efficient JSON encoding/decoding
- **Middleware**: Request logging and recovery

## Configuration

Environment variables:
- `DATABASE_URL`: PostgreSQL connection string
- `REDIS_URL`: Redis connection string
- `ENVIRONMENT`: Application environment (development/production)
- `PORT`: HTTP server port (default: 8080)

## Monitoring & Health Checks

- **Health Endpoint**: `/health` for service health
- **Database Health**: Connection status check
- **Cache Health**: Redis connection status
- **Structured Logging**: JSON formatted logs