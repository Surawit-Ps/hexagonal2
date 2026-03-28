package handler

import (
	"hexagonal2/core/entity"
	"hexagonal2/core/ports"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userSer ports.UserServices
}

func NewUserHandler(userSer ports.UserServices) userHandler {
	return userHandler{userSer: userSer}
}

func (h userHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userSer.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(newResponse(false, "Internal Server Error", nil))
	}
	return c.JSON(newResponse(true, "Success", users))
}

func (h userHandler) GetAUser(c *fiber.Ctx) error {
	user, err := h.userSer.GetUser(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(newResponse(false, "User not found", nil))
	}
	return c.JSON(newResponse(true, "Success", user))
}

func (h userHandler) AddUser(c *fiber.Ctx) error {
	var person entity.User
	if err := c.BodyParser(&person); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(newResponse(false, "Invalid request body", nil))
	}
	if err := h.userSer.AddUser(person); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(newResponse(false, "Internal Server Error", nil))
	}
	return c.JSON(newResponse(true, "Success", nil))
}
