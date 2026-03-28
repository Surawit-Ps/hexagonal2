package handler

import (
	"hexagonal2/core/entity"
	"hexagonal2/core/ports"

	"github.com/gofiber/fiber/v2"
)

type dogHandler struct{
	dogSer ports.DogServices
}

func NewDogHandler(dogSer ports.DogServices)dogHandler{
	return dogHandler{dogSer: dogSer}
}

func(h dogHandler)GetAllDogs(c *fiber.Ctx)error{
	dog ,err:= h.dogSer.GetAllDogs()
	if  err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(newResponse(false, "Internal Server Error", nil))
	}
	return c.JSON(newResponse(true, "Success", dog))
}

func(h dogHandler)GetADogs(c *fiber.Ctx)error{
	dog,err:=h.dogSer.GetDog(c.Params("id"))
	if err!=nil{
		return c.JSON(fiber.ErrNotFound)
	}
	return c.Status(fiber.StatusOK).JSON(newResponse(true, "Success", dog))
}

func(h dogHandler)AddDog(c *fiber.Ctx)error{
	var dog entity.Dogs
	err := c.BodyParser(&dog)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(newResponse(false, "Invalid request body", nil))
	}
	id := dog.UserID
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(newResponse(false, "Human ID is required", nil))
	}
	if err := h.dogSer.AddDog(dog, id); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(newResponse(false, "Internal Server Error", nil))
	}
	return c.Status(fiber.StatusCreated).JSON(newResponse(true, "Success", nil))
}