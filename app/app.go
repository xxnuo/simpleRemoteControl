package main

import (
	"fmt"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/xxnuo/simpleRemoteControl/internal/api"
	"github.com/xxnuo/simpleRemoteControl/internal/log"
)

// Cobra config loaded
func App() {

	// 日志初始化
	Logger = log.NewLogger(Cfg.IsJsonLog)
	Logger.Info().Msg("Starting app")
	Logger.Info().Msgf("Config: %v", Cfg)

	// API 服务器初始化
	ApiServer = fiber.New()
	ApiServer.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &Logger,
	}))

	ApiServer.Use(cors.New())
	api.SetupRouter(ApiServer)

	// 启动 API 服务器
	if err := ApiServer.Listen(fmt.Sprintf("%s:%d", Cfg.Addr, Cfg.Port)); err != nil {
		Logger.Fatal().Err(err).Msg("Fiber app error")
	}
}
