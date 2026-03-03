# Go Backend Fundamentals: HTTP & CRUD

This repository contains my progress in learning backend engineering with Go. It follows a practical implementation of a RESTful API, covering the core concepts of the HTTP protocol, routing, and data management.

## 🚀 Learning Objectives
The goal of this project was to understand how a server communicates with a client. Key concepts covered include:
* **HTTP Basics:** Understanding Request/Response cycles, Headers, and Bodies.
* **Status Codes:** Implementing `200 OK`, `201 Created`, `400 Bad Request`, `404 Not Found`, and `500 Internal Server Error`.
* **Standard Library:** Using Go's `net/http` package to build a web server.
* **External Routers:** Implementing `gorilla/mux` for advanced path parameters.
* **CRUD Pattern:** Creating, Reading, Updating, and Deleting resources (Users) in memory.
* **JSON Processing:** Marshalling and Unmarshalling data for API communication.

## 🛠️ Project Structure
* `main.go` - The entry point of the application.
* `pkg/server/` - Organized server logic and route handlers.
* `go.mod` - Dependency management.

## 🛠️ Setup & Installation
1. **Clone the repository:**
   ```bash
   git clone [https://github.com/Yogeshpant992/](https://github.com/Yogeshpant992/http-basic-with-go).git
   cd http-basic-with-go
