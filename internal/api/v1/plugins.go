package v1

import (
	"github.com/gofiber/fiber/v2"
)

func RunPlugin(c *fiber.Ctx) error {
	c.Context().Logger().Printf("Running plugin...")
	return c.JSON(fiber.Map{"status": "success", "message": "Running plugin...", "data": nil})
}
