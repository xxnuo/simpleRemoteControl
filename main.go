package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
			Logger.Info().Msg("config called")
		},
	}
	setCmd = &cobra.Command{
		Use:   "set",
		Short: "保存新配置",
		Long:  `修改指定配置项目并保存`,
		Run: func(cmd *cobra.Command, args []string) {
			Logger.Info().Msg("set called")
		},
	}
	clearCmd = &cobra.Command{
		Use:   "clear",
		Short: "清除配置",
		Long:  `清除所有或指定配置项`,
		Run: func(cmd *cobra.Command, args []string) {
			Logger.Info().Msg("clear called")
		},
	}
)

func initCfg() {
	if Cfg.File != "" {
		viper.SetConfigFile(Cfg.File)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".rc")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "使用配置文件:", viper.ConfigFileUsed())
	}
}

func main() {
	cobra.OnInitialize(initCfg)

	rootCmd.PersistentFlags().StringVarP(&Cfg.File, "config", "c", "", "配置文件路径 (默认 $HOME/.rc.yaml)")
	rootCmd.PersistentFlags().StringVarP(&Cfg.Addr, "bind", "b", "0.0.0.0", "服务器绑定地址")
	rootCmd.PersistentFlags().IntVarP(&Cfg.Port, "port", "p", 10101, "服务器绑定端口")
	rootCmd.PersistentFlags().BoolVarP(&Cfg.IsJsonLog, "json", "j", false, "以 JSON 格式输出日志")
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(setCmd)
	configCmd.AddCommand(clearCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
