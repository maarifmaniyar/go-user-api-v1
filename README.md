# Go User API – DOB & Age Calculation
📌 Overview

This project is a RESTful backend API built using Go and GoFiber to manage users with their Name and Date of Birth (DOB).
The API dynamically calculates and returns the user’s age when fetching user details.

The project follows a clean, layered architecture and uses PostgreSQL with SQLC for database access.

Important Note: For Password
## 🔐 Database Credentials

For security reasons, database credentials are **not included** in this repository.

If you need access, please contact me:

📧 **Email:** maarifmaniyar5@gmail.com

# 👨‍💻 Author

Mohammed Maarif Maniyar

B.Tech – Artificial Intelligence & Machine Learning

Passionate about Backend Development & System Design

📧 Email: maarifmaniyar5@gmail.com

🔗 Portfolio: (https://maarifmaniyar.github.io/)


#  Tech Stack

Go

GoFiber

PostgreSQL

SQLC

Thunder Client / Postman

pgAdmin

lib/pq (PostgreSQL driver)

📂 Project Structure
go-user-api-v1/
├── cmd/
│   └── server/
│       └── main.go
├── db/
│   ├── migrations/
│   │   └── schema.sql
│   └── sqlc/
│       ├── queries.sql
│       └── generated files
├── internal/
│   ├── handler/
│   ├── repository/
│   ├── routes/
│   ├── service/
│   ├── middleware/
│   ├── models/
│   └── logger/
├── go.mod
└── README.md

#  Database Schema
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  dob DATE NOT NULL
);

#  Setup Instructions
1️⃣ Prerequisites

Go (v1.20+ recommended)

PostgreSQL

pgAdmin

Thunder Client or Postman

2️⃣ Database Setup

Open pgAdmin

Connect to PostgreSQL

Select database: postgres

Open Query Tool

Run:

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  dob DATE NOT NULL
);

3️⃣ Update Database Credentials

Edit cmd/server/main.go:

db, err := sql.Open(
	"postgres",
	"postgres://postgres:YOUR_PASSWORD@localhost:5432/postgres?sslmode=disable",
)


#  Replace YOUR_PASSWORD with your PostgreSQL password.

4️⃣ Install Dependencies
go mod tidy

5️⃣ Run the Server
go run cmd/server/main.go


You should see:

Server running on :8080

🔌 API Endpoints
➕ Create User

POST /users

Request Body
{
  "name": "Alice",
  "dob": "1990-05-10"
}

Response
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}

🔍 Get User by ID

GET /users/{id}

Response
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}


📌 Age is calculated dynamically using Go’s time package.

#  Key Features

Clean layered architecture

Dynamic age calculation

SQLC for type-safe database queries

PostgreSQL integration

RESTful API design

Error handling and logging

#  How to Test

Use Thunder Client or Postman:

Start the server

Send POST request to /users

Fetch data using GET /users/{id}

# Notes

Age is not stored in the database

Age is calculated dynamically when fetching user details

Database connection issues usually result from incorrect credentials

## 👨‍💻 Author

Mohammed Maarif Maniyar

B.Tech – Artificial Intelligence & Machine Learning

Passionate about Backend Development & System Design

📧 Email: maarifmaniyar5@gmail.com

🔗 Portfolio: (https://maarifmaniyar.github.io/)

# 🏁 Conclusion

This project demonstrates a complete backend API workflow using Go, PostgreSQL, and SQLC.

# ⚠️ Disclaimer

This project is developed strictly for educational, learning, and evaluation purposes.

It is not intended for production use.

The author is not responsible for any misuse of this code outside its intended purpose.

