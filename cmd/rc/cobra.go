package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xxnuo/simpleRemoteControl/internal/v"
)

func initCfg() {

	// 日志初始化
	consoleWriter := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime})
	if v.Cfg.IsJsonLog {
		consoleWriter = zerolog.New(os.Stderr)
	}
	v.Logger = consoleWriter.With().Timestamp().Logger()

	v.Logger.Info().Msg("Starting app")
	v.Logger.Info().Msgf("Config: %v", v.Cfg)

	if v.Cfg.IsDebug {
		_cd, err := os.Getwd()
		v.CheckErr(err)
		v.Cfg.File = filepath.Join(_cd, "rc.yaml")
	}

	if v.Cfg.File != "" {
		viper.SetConfigFile(v.Cfg.File)
		v.Cfg.WorkDir = filepath.Dir(v.Cfg.File)
	} else {
		home, err := os.UserHomeDir()
		v.CheckErr(err)
		v.Cfg.WorkDir = filepath.Join(home, ".config", "rc")
		v.Cfg.File = filepath.Join(v.Cfg.WorkDir, "rc.yaml")
		viper.SetConfigFile(v.Cfg.File)
	}

	v.Cfg.PluginsDir = filepath.Join(v.Cfg.WorkDir, "plugins")

	v.Logger.Info().Msgf("Loading working directory: %s", v.Cfg.WorkDir)
	v.Logger.Info().Msgf("Loading config file: %s", v.Cfg.File)

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		v.Logger.Info().Msgf("Current config file: %s", viper.ConfigFileUsed())
	}
}

func main() {

	rootCmd.PersistentFlags().StringVarP(&v.Cfg.File, "config", "c", "", "配置文件路径 (默认 $HOME/.config/rc/rc.yaml)")
	rootCmd.PersistentFlags().StringVarP(&v.Cfg.Addr, "bind", "b", "0.0.0.0", "服务器绑定地址")
	rootCmd.PersistentFlags().IntVarP(&v.Cfg.Port, "port", "p", 10101, "服务器绑定端口")
	rootCmd.PersistentFlags().BoolVarP(&v.Cfg.IsJsonLog, "json", "j", false, "以 JSON 格式输出日志")
	// rootCmd.PersistentFlags().BoolVarP(&v.Cfg.IsDebug, "debug", "d", false, "以当前目录为工作目录, 配置文件路径为当前目录下的 rc.yaml")
	rootCmd.PersistentFlags().BoolVarP(&v.Cfg.IsDebug, "debug", "d", true, "以当前目录为工作目录, 配置文件路径为当前目录下的 rc.yaml")
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(setCmd)
	configCmd.AddCommand(clearCmd)

	// println("before")
	cobra.OnInitialize(initCfg)
	err := rootCmd.Execute()
	// println("after")
	if err != nil {
		os.Exit(1)
	}
}

var (
	rootCmd = &cobra.Command{
		Use:     "rc",
		Short:   "远程控制工具",
		Long:    `通过网页管理的个人 PC 远程控制工具`,
		Version: "v0.0.1",
		// App 主函数启动服务器
		Run: func(cmd *cobra.Command, args []string) { App() },
	}
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "修改持久设置",
		Long:  `通过此命令创建和修改配置并保存到配置文件 (默认位置 $HOME/.rc.yaml)`,
		Run: func(cmd *cobra.Command, args []string) {
			v.Logger.Info().Msg("config called")
		},
	}
	setCmd = &cobra.Command{
		Use:   "set",
		Short: "保存新配置",
		Long:  `修改指定配置项目并保存`,
		Run: func(cmd *cobra.Command, args []string) {
			v.Logger.Info().Msg("set called")
		},
	}
	clearCmd = &cobra.Command{
		Use:   "clear",
		Short: "清除配置",
		Long:  `清除所有或指定配置项`,
		Run: func(cmd *cobra.Command, args []string) {
			v.Logger.Info().Msg("clear called")
		},
	}
)
