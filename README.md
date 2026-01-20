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
└── routes/          # HTTP routes
└── middlewares/     # Project middlewares (auth filter)

pkg/
├── auth/            # Authentication utilities (generate and verify tokens)
├── db/              # Database connection (if any)
```

## 🛠️ Technologies

- **Go 1.25**
- **Gin** - Web framework
- **Layered Architecture** - Separation of concerns
- **Swagger** - Docs

## 🚀 How to Run

### Prerequisites
- Go 1.25

### Installation

1. Clone the repository:
```bash
git clone https://github.com/PedroVidalDev/go-product-api
cd go-product-api
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the project:
```bash
go run cmd/api/main.go
```

4. If you want to run using Docker (remember to define the env variables in a .env file):
```bash
docker-compose up --build
```

The API will be available at `http://localhost:8080`

## 📍 API Routes

### Authentication Routes

#### Register User
- **POST** `/register`
- **Description**: Creates a new user account
- **Body**:
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword"
}
```
- **Response**: User object (201 Created)

#### Login
- **POST** `/login`
- **Description**: Authenticates a user and returns a JWT token
- **Body**:
```json
{
  "email": "john@example.com",
  "password": "securepassword"
}
```
- **Response**: JWT token (200 OK)

### Product Routes (Protected)

All product routes require authentication via Bearer token in the Authorization header.

#### Get All Products
- **GET** `/products`
- **Description**: Retrieves all products with their associated users
- **Headers**: `Authorization: Bearer <token>`
- **Response**: Array of products (200 OK)

#### Create Product
- **POST** `/products`
- **Description**: Creates a new product linked to the authenticated user
- **Headers**: `Authorization: Bearer <token>`
- **Body**:
```json
{
  "name": "Product Name",
  "price": 99.99
}
```
- **Response**: Created product object (201 Created)

### Documentation

- **Swagger UI**: `http://localhost:8080/swagger/index.html`
- Interactive API documentation with request/response examples

## 📖 Concepts Learned

- **Structs**: Custom data structures
- **Interfaces**: Implicit contracts in Go
- **Receivers**: Methods associated with structs
- **Pointers**: Memory management with `*` and `&`
- **Multiple returns**: Functions returning values and errors
- **Gin Context**: HTTP request handling
- **JSON Binding**: Automatic serialization/deserialization
- **Database Operations**: CRUD operations
- **Entity Relationships**: Managing related data
- **Dependency Injection**: Manual dependency injection
- **Swagger Integration**: API documentation
- **Modular Code Organization**: Clean project structure

## 🎓 Study Resources

- [Official Go Documentation](https://go.dev/doc/)

## 📝 Notes

- This is an educational project, not optimized for production
- Focused on learning Go fundamentals
