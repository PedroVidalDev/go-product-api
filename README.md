# PRODUCT API - GO

## 📚 About the Project

This is a **study project** developed to learn the **Go (Golang)** programming language by building a simple REST API for product management.

## 🎯 Learning Goals

- Go project structure
- Building REST APIs with Gin Framework
- Implementing architectural patterns (Repository, Service, Controller)
- Working with structs, interfaces, and pointers
- JSON manipulation
- Dependency management with Go Modules

## 🏗️ Architecture

The project follows a layered architecture:

```
internal/
├── controllers/     # HTTP handlers (receives requests)
├── services/        # Business logic
├── repositories/    # Data access (in-memory)
├── models/          # Data structures
└── dtos/            # Data Transfer Objects
```

## 🛠️ Technologies

- **Go 1.21+**
- **Gin** - Web framework
- **Layered Architecture** - Separation of concerns

## 🚀 How to Run

### Prerequisites
- Go 1.21 or higher installed

### Installation

1. Clone the repository:
```bash
git clone <your-repository>
cd go-product-api
```

2. Install dependencies:
```bash
go mod download
```

3. Run the project:
```bash
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

## 📖 Concepts Learned

- **Structs**: Custom data structures
- **Interfaces**: Implicit contracts in Go
- **Receivers**: Methods associated with structs
- **Pointers**: Memory management with `*` and `&`
- **Multiple returns**: Functions returning values and errors
- **Gin Context**: HTTP request handling
- **JSON Binding**: Automatic serialization/deserialization
- **Dependency Injection**: Manual dependency injection

## 🎓 Study Resources

- [Official Go Documentation](https://go.dev/doc/)

## 📝 Notes

- This is an educational project, not optimized for production
- Focused on learning Go fundamentals
