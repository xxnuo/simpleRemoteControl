package engine

import (
	"github.com/rs/zerolog"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// 引擎加载插件
func Load(pluginsDir string, l zerolog.Logger) Engine {
	// 创建 infoWriter 和 errorWriter
	infoWriter := LoggerWriter{logger: l, level: zerolog.InfoLevel}
	errorWriter := LoggerWriter{logger: l, level: zerolog.ErrorLevel}

	i := interp.New(interp.Options{
		GoPath: pluginsDir,
		Stdout: infoWriter,
		Stderr: errorWriter,
	})

	i.Use(stdlib.Symbols)
	return Engine{i: i}
}
