package v

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	"github.com/xxnuo/simpleRemoteControl/internal/engine"
)

// 命令行参数集合
type CobraConfig struct {
	File       string // 配置文件路径
	IsJsonLog  bool   // 是否输出 JSON 格式日志
	IsDebug    bool   // 是否开启调试模式
	Addr       string // 监听地址 默认 0.0.0.0
	Port       int    // 监听端口 默认 10101
	WorkDir    string // 工作目录
	PluginsDir string // 插件目录, 默认为工作目录下的 plugins 目录
}

// 全局变量
var (
	Cfg    CobraConfig    // 命令行参数
	Logger zerolog.Logger // 日志对象

	PluginEngine  engine.Engine         // 插件引擎对象
	PluginHandles []engine.PluginHandle // 插件处理函数集合

	ApiServer     *fiber.App   // API 服务器对象
	RootRouter    fiber.Router // 根路由对象
	PluginsRouter fiber.Router // 插件路由对象

)

// CheckErr prints the msg with the prefix 'Error:' and exits with error code 1. If the msg is nil, it does nothing.
func CheckErr(msg error) {
	if msg != nil {
		Logger.Error().Err(msg)
		os.Exit(1)
	}
}
