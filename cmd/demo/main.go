package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/rs/zerolog"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// levelWriter 结构体，支持通用日志级别
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

func main() {
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	l.Info().Msg("Starting plugin")

	// 创建 infoWriter 和 errorWriter
	infoWriter := LoggerWriter{logger: l, level: zerolog.InfoLevel}
	errorWriter := LoggerWriter{logger: l, level: zerolog.ErrorLevel}

	i := interp.New(interp.Options{
		Stdout: infoWriter,
		Stderr: errorWriter,
	})
	i.Use(stdlib.Symbols)
	_, filename, _, _ := runtime.Caller(0)

	s, e := os.ReadFile(filepath.Join(filepath.Dir(filename), "plugin1.go"))
	if e != nil {
		panic(e)
	}
	src := string(s)

	_, err := i.Eval(src)
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("plugin1.Init")
	if err != nil {
		panic(err)
	}

	v.Call(nil)
}
