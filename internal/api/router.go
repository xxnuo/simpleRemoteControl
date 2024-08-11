package api

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/xxnuo/simpleRemoteControl/internal/api/v1"
)

func SetupRouter(a *fiber.App) {
	root := a.Group("/api/v1")
	root.Get("/", v1.Hello)

	// Auth
	auth := root.Group("/auth")
	auth.Post("/login", v1.Auth)
}
