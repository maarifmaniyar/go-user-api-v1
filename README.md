# Go User API 

A simple RESTful User Management API built using **Go**, **Fiber**, **PostgreSQL**, and **sqlc**.

This project demonstrates clean architecture with handlers, routes, repositories, and database access using generated SQL code.

---

#  Features

- Create a user
- Get user by ID
- PostgreSQL database
- sqlc for type-safe SQL
- Fiber web framework
- Thunder Client / Postman tested

---

#  Tech Stack

- **Go**
- **Fiber**
- **PostgreSQL**
- **sqlc**
- **pgAdmin**
- **Thunder Client / Postman**

---

#   Project Structure

go-user-api-v1/
│
├── cmd/
│ └── server/
│ └── main.go
│
├── internal/
│ ├── handler/
│ ├── repository/
│ ├── routes/
│
├── db/
│ └── sqlc/
│
├── migrations/
│ ├── schema.sql
│ └── queries.sql
│
├── go.mod
├── go.sum
└── README.md


---

##  Database Schema

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);

# How to Run the Project
1️. Start PostgreSQL

Make sure PostgreSQL is running and the database exists:
CREATE DATABASE userdb;

2️. Update DB Connection (main.go)
postgres://postgres:YOUR_PASSWORD@localhost:5432/userdb?sslmode=disable

3️. Run the Server
go run cmd/server/main.go


Server will start on:

http://localhost:8080

 API Endpoints
 Create User
POST /users


Request Body

{
  "name": "Alice",
  "dob": "1990-05-10"
}

-> Get User by ID
GET /users/:id


Example:

GET /users/1

-> Testing

Use Thunder Client or Postman

Content-Type: application/json

-> Future Improvements

Get all users

- Update & delete user

- Environment variables

- Docker support

- Authentication

- Swagger documentation

👤 Author

MOhammed Maarif Maniyar
GitHub: https://github.com/maarifmaniyar




