package handler

import (
	"hgo/core/entity"
	"hgo/core/ports"

	"github.com/gofiber/fiber/v2"
)

type humanHandler struct {
	humanSer ports.HumanServices
}

func NewHumanHandler(humanSer ports.HumanServices) humanHandler {
	return humanHandler{humanSer: humanSer}
}

func (h humanHandler) GetAllUsers(c *fiber.Ctx) error {
	humans, err := h.humanSer.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(newResponse(false, "Internal Server Error", nil))
	}
	return c.JSON(newResponse(true, "Success", humans))
}

func (h humanHandler) GetAUser(c *fiber.Ctx) error {
	human, err := h.humanSer.GetUser(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(newResponse(false, "User not found", nil))
	}
	return c.JSON(newResponse(true, "Success", human))
}

func (h humanHandler) AddUser(c *fiber.Ctx) error {
	var person entity.Humans
	if err := c.BodyParser(&person); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(newResponse(false, "Invalid request body", nil))
	}
	if err := h.humanSer.AddUser(person); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(newResponse(false, "Internal Server Error", nil))
	}
	return c.JSON(newResponse(true, "Success", nil))
}
