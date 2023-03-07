package controllers

import (
	"github.com/abe27/cvst20/api/models"
	"github.com/gofiber/fiber/v2"
)

func HandlerHello(c *fiber.Ctx) error {
	var r models.Response
	r.Message = "Test Message,Hello World!"
	return c.Status(fiber.StatusOK).JSON(&r)
}
