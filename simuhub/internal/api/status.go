package api

import "github.com/gofiber/fiber/v3"

func HandleGetStatus(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
