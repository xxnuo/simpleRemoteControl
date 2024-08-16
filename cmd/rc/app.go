package main

import (
	"fmt"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	v1 "github.com/xxnuo/simpleRemoteControl/internal/api/v1"
	"github.com/xxnuo/simpleRemoteControl/internal/engine"
	"github.com/xxnuo/simpleRemoteControl/internal/v"
)

// Cobra config loaded
func App() {

	// 加载插件
	v.Logger.Info().Msg("Loading plugins...")
	LoadPlugins()

	// API 服务器初始化
	v.ApiServer = fiber.New()
	v.ApiServer.Use(fiberzerolog.New(fiberzerolog.Config{
		Fields: []string{"ip", "status", "method", "url", "error"},
		Logger: &v.Logger,
	}))

	v.ApiServer.Use(cors.New())
	v.RootRouter = InitRouter(v.ApiServer)
	v.PluginsRouter = RegisterPluginsRouter(v.RootRouter, v.PluginHandles)

	// 启动 API 服务器
	// PrintRoutes(v.ApiServer)

	if err := v.ApiServer.Listen(fmt.Sprintf("%s:%d", v.Cfg.Addr, v.Cfg.Port)); err != nil {
		v.Logger.Fatal().Err(err).Msg("Fiber app error")
	}
}

// 读取 plugins 目录下所有插件并加载
func LoadPlugins() {
	v.Logger.Info().Msg("Reloading plugins...")

	v.PluginEngine = engine.New(v.Cfg.PluginsDir, v.Logger)
	v.PluginHandles = v.PluginEngine.LoadAll(v.Cfg.PluginsDir)
}

// 初始化路由
func InitRouter(a *fiber.App) fiber.Router {
	root := a.Group("/api/v1")
	root.Get("/", v1.Hello)

	token := root.Group("/token")
	token.Post("/new", v1.NewToken)

	root.Post("/reload", ReloadPluginsRouter)

	return root
}

// 初次注册所有插件路由
func RegisterPluginsRouter(rootRouter fiber.Router, hs []engine.PluginHandle) fiber.Router {
	getPackageNames := func(hs []engine.PluginHandle) []string {
		names := make([]string, len(hs))
		for i, h := range hs {
			names[i] = h.PackageName
		}
		return names
	}

	plugins := rootRouter.Group("/plugins")
	plugins.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(getPackageNames(hs))
	})
	for _, h := range hs {
		RegisterPluginRouter(plugins, h)
	}
	return plugins
}

// 注册单个插件路由
func RegisterPluginRouter(pluginsRouter fiber.Router, e engine.PluginHandle) {
	pluginsRouter.Post("/"+e.PackageName, func(c *fiber.Ctx) error {
		msg, err := e.Run(string(c.Body()))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(msg)
		}
		return c.Status(fiber.StatusOK).JSON(msg)
	})
}

// 重新加载插件及插件的路由
func ReloadPluginsRouter(c *fiber.Ctx) error {
	LoadPlugins()

	v.PluginsRouter = RegisterPluginsRouter(v.RootRouter, v.PluginHandles)
	return c.SendString("Plugins reloaded")
}

// 打印所有路由信息
func PrintRoutes(a *fiber.App) {
	routes := a.GetRoutes()
	v.Logger.Debug().Msg("Routes:")
	for _, route := range routes {
		v.Logger.Debug().Msgf("Method: %s, Path: %s\n", route.Method, route.Path)
	}
}
