package engine

import (
	"github.com/rs/zerolog"
	"github.com/traefik/yaegi/interp"
)

type Engine struct {
	Name string
	i    *interp.Interpreter
}

type Plugin struct {
	Name string
}

// LoggerWriter 实现 io.Writer 接口，支持通用日志级别，重定向插件输出
type LoggerWriter struct {
	logger zerolog.Logger
	level  zerolog.Level
}

func (w LoggerWriter) Write(p []byte) (n int, err error) {
	// 根据指定的日志级别记录消息
	switch w.level {
	case zerolog.InfoLevel:
		w.logger.Info().Msg(string(p))
	case zerolog.ErrorLevel:
		w.logger.Error().Msg(string(p))
	}
	return len(p), nil
}
