package main

import (
	"github.com/gofiber/fiber/v2"
)

func ImageController(c *fiber.Ctx) error {

	var file []byte
	value := c.Query("value")

	option := c.Params("option")

	if value == "" {
		return c.Status(fiber.StatusBadRequest).SendString("El parámetro 'nombre' es obligatorio")
	}

	switch option {
	case "avatar":

		file = CreateAvatar(value)

		c.Set("Content-Type", "image/svg+xml")

	case "logo":

		file = CreateInitial(value)

		c.Set("Content-Type", "image/svg+xml")

	}

	return c.Send(file)

}
