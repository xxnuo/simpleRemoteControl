package v

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/xxnuo/simpleRemoteControl/internal/engine"
	"github.com/xxnuo/simpleRemoteControl/internal/log"
)

// 命令行参数集合
type CobraConfig struct {
	Addr      string
	Port      int
	WorkDir   string
	File      string
	IsJsonLog bool
	IsDebug   bool
}

// 全局变量
var (
	Cfg          CobraConfig   // 命令行参数
	Logger       log.Logger    // 日志对象
	ApiServer    *fiber.App    // API 服务器对象
	PluginEngine engine.Engine // 插件引擎对象
)

// CheckErr prints the msg with the prefix 'Error:' and exits with error code 1. If the msg is nil, it does nothing.
func CheckErr(msg error) {
	if msg != nil {
		Logger.Error().Err(msg)
		os.Exit(1)
	}
}
