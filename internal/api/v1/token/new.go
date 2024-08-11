package v1

import "github.com/gofiber/fiber/v2"

func NewToken(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Success auth", "data": nil})

}
