package handler

import (
	"log"
	"strconv"
	"time"

	"go-user-api-v1/internal/repository"
	"go-user-api-v1/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Dob  string `json:"dob"`
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.repo.CreateUser(c.Context(), req.Name, dob)
	if err != nil {
		log.Println("CreateUser error:", err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
	})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user, err := h.repo.GetUserByID(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrNotFound
	}

	age := service.CalculateAge(user.Dob)

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
		"age":  age,
	})
}
