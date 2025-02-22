# Freezer Inventory API

A RESTful API service for managing freezer inventory, built with Go and PostgreSQL.

## Features

- 🔐 JWT Authentication & Authorization
- 📦 Item Management with Categories and Tags
- 📊 Inventory Tracking
- 👥 User Management with Role-based Access
- 🔍 Search Functionality
- 📝 Detailed Activity Logging
- ⚡ Rate Limiting
- 🔄 CORS Support

## Tech Stack

- Go 1.22+
- PostgreSQL
- GORM (ORM)
- Gin (Web Framework)
- JWT for Authentication

## Prerequisites

- Go 1.22 or higher
- PostgreSQL 12 or higher
- Make (optional, for using Makefile commands)

## Quick Start

1. Clone the repository:
```bash
git clone https://github.com/gevgev/freezer-inventory.git
cd freezer-inventory
```

2. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your database credentials and JWT secret
```

3. Install dependencies:
```bash
go mod download
```

4. Run database migrations:
```bash
go run cmd/migrate/main.go up
```

5. Start the server:
```bash
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

## API Documentation

See [API.md](API.md) for detailed API documentation.

## Development

### Project Structure
```
.
├── cmd/                    # Application entry points
│   ├── api/               # API server
│   └── migrate/           # Database migrations
├── internal/              # Private application code
│   ├── api/              # API layer
│   │   ├── handlers/     # Request handlers
│   │   ├── middleware/   # HTTP middleware
│   │   └── router/       # Route definitions
│   ├── models/           # Data models
│   ├── repository/       # Data access layer
│   └── service/          # Business logic
├── migrations/           # SQL migrations
└── pkg/                  # Public libraries
```

### Running Tests
```bash
go test ./...
```

### Available Make Commands
```bash
make build      # Build the application
make run        # Run the application
make test       # Run tests
make migrate    # Run database migrations
make docker     # Build Docker image
```

## Testing with Postman

1. Import the Postman collection from `Freezer_Inventory.postman_collection.json`
2. Import the environment from `Freezer_Inventory.postman_environment.json`
3. Set your `base_url` in the environment (default: `http://localhost:8080`)
4. Use the collection's pre-request scripts to handle authentication automatically

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

Gev Gev - [@gevgev](https://github.com/gevgev)

Project Link: [https://github.com/gevgev/freezer-inventory](https://github.com/gevgev/freezer-inventory) 