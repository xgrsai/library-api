# 📚 Book Library API (Go, net/http)

A lightweight REST API for managing a book library, built in **pure Go** using only the standard library (`net/http`).  
The project focuses on understanding core backend concepts such as routing, middleware, concurrency, and authentication without relying on external frameworks.

---

## 🚀 Features

- User authentication (register/login)
- JWT-based protected routes
- CRUD operations for books
- Public and private endpoints
- In-memory storage with concurrency safety (`sync.RWMutex`)
- Custom routing using `net/http`
- Middleware for authentication
- Background worker using goroutines for event logging
- Input validation
- Docker support

---

## 📌 API Endpoints

### Authentication
```

POST /auth/register
POST /auth/login

```

### Books (Public)
```

GET    /books
GET    /books/{id}

```

### Books (Protected - JWT required)
```

POST   /books
PUT    /books/{id}
DELETE /books/{id}

````

---

## 🧠 Architecture Overview

- **Structs** act as data models (`Book`, `User`, `Response`)
- **net/http** handles all routing manually
- **Middleware** validates JWT tokens
- **Goroutines + channels** process background events
- **sync.RWMutex** ensures safe concurrent access to in-memory storage

---

## 🛠 Tech Stack

- Go (standard library only)
- net/http
- encoding/json
- sync (mutexes)
- goroutines & channels
- JWT (custom implementation or minimal dependency if added)
- Docker

---

## 🐳 Running with Docker

### Build and run:
```bash
docker build -t book-library-api .
docker run -p 8080:8080 book-library-api
````

---

## 📂 Project Goals

This project is designed to practice:

* Building APIs without frameworks
* Understanding Go concurrency model
* Writing middleware manually
* Structuring backend services in Go
* Working with HTTP at a low level

---

## 📈 Possible Improvements

* Replace in-memory storage with PostgreSQL
* Add Redis caching
* Add Swagger documentation
* Add rate limiting middleware
* Add unit and integration tests
* Improve JWT security & refresh token
