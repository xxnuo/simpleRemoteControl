package api

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/xxnuo/simpleRemoteControl/internal/api/v1"
)

func InitRouter(a *fiber.App) {
	root := a.Group("/api/v1")
	root.Get("/", v1.Hello)

	token := root.Group("/token")
	token.Post("/new", v1.NewToken)
}

func UpdateRouter(a *fiber.App) {}
