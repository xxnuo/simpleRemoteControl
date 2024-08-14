package engine

import "github.com/rs/zerolog"

// loggerWriter 实现 io.Writer 接口，支持通用日志级别，重定向插件输出
type loggerWriter struct {
	logger zerolog.Logger
	level  zerolog.Level
}

func (w loggerWriter) Write(p []byte) (n int, err error) {
	// 根据指定的日志级别记录消息
	switch w.level {
	case zerolog.InfoLevel:
		w.logger.Info().Msg(string(p))
	case zerolog.ErrorLevel:
		w.logger.Error().Msg(string(p))
	}
	return len(p), nil
}
