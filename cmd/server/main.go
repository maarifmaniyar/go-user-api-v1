package main

import (
	"database/sql"
	"log"

	"go-user-api-v1/db/sqlc"
	"go-user-api-v1/internal/handler"
	"go-user-api-v1/internal/repository"
	"go-user-api-v1/internal/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	// ⚠️ Update DB credentials if needed
	db, err := sql.Open(
		"postgres",
		"postgres://postgres:postgres123@localhost:5432/postgres?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	queries := sqlc.New(db)
	repo := repository.NewUserRepository(queries)
	userHandler := handler.NewUserHandler(repo)

	app := fiber.New()
	routes.Register(app, userHandler)

	log.Println("Server running on :8080")
	log.Fatal(app.Listen(":8080"))
}
