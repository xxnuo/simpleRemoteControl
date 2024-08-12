package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xxnuo/simpleRemoteControl/internal/v"
)

func Hello(c *fiber.Ctx) error {
	v.Logger.Info().Msg("Hello to console")
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
