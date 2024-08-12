package main

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/xxnuo/simpleRemoteControl/internal/api"
	"github.com/xxnuo/simpleRemoteControl/internal/v"
)

// Cobra config loaded
func App() {

	// API 服务器初始化
	v.ApiServer = fiber.New()
	v.ApiServer.Use(fiberzerolog.New(fiberzerolog.Config{
		Fields: []string{"ip", "status", "method", "url", "error"},
		Logger: &v.Logger,
	}))

	v.ApiServer.Use(cors.New())
	api.InitRouter(v.ApiServer)

	// 加载插件

	// 启动 API 服务器
	// if err := v.ApiServer.Listen(fmt.Sprintf("%s:%d", v.Cfg.Addr, v.Cfg.Port)); err != nil {
	// 	v.Logger.Fatal().Err(err).Msg("Fiber app error")
	// }
}
