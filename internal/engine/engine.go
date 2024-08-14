package engine

import (
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"github.com/traefik/yaegi/stdlib/syscall"
	"github.com/traefik/yaegi/stdlib/unrestricted"
	"github.com/traefik/yaegi/stdlib/unsafe"
	"github.com/xxnuo/simpleRemoteControl/internal/tool"
)

type Engine struct {
	i *interp.Interpreter
	l zerolog.Logger
}

type PluginHandle struct {
	PackageName string
	Run         func([]byte) (msg []byte, err error)
}

// 创建 Engine 实例
// pluginsDir: 插件目录，用于指定插件的加载路径和 GOPATH
// l: 日志实例
func New(GOPATH string, l zerolog.Logger) Engine {
	// 创建 infoWriter 和 errorWriter
	infoWriter := loggerWriter{logger: l, level: zerolog.InfoLevel}
	errorWriter := loggerWriter{logger: l, level: zerolog.ErrorLevel}

	i := interp.New(interp.Options{
		GoPath: GOPATH,
		Stdout: infoWriter,
		Stderr: errorWriter,
	})

	i.Use(stdlib.Symbols)
	i.Use(syscall.Symbols)
	i.Use(unsafe.Symbols)
	i.Use(unrestricted.Symbols)

	return Engine{i: i, l: l}
}

// 加载单个插件文件
// dirPath: 插件目录路径，默认插件目录名为插件主包名，包名.Run() 对应函数即为插件入口函数
// 插件目录结构示例：
//   - plugin1/ // 则插件包名必须为 plugin1，必须存在包名.Run() 对应函数即为插件入口函数
//     -- plugin1.go // 任意功能
//     -- cfg.go // 任意功能
func (e *Engine) Load(dirPath string) PluginHandle {

	_path, err := filepath.Abs(dirPath)
	if err != nil {
		e.l.Error().Err(err).Msgf("Failed to get full path of %s", dirPath)
		return PluginHandle{}
	}
	dirName := filepath.Base(_path)

	// 遍历插件目录路径下所有文件，加载go文件
	files, err := filepath.Glob(filepath.Join(_path, "*.go"))
	if err != nil {
		e.l.Error().Err(err).Msgf("Failed to glob %s", filepath.Join(_path, "*.go"))
		return PluginHandle{}
	}

	for _, file := range files {
		e.l.Info().Msgf("Loading file %s", file)
		_, err := e.i.EvalPath(file)
		if err != nil {
			e.l.Error().Err(err).Msgf("Failed to load file %s", file)
		}
	}

	// 加载插件入口函数
	v, err := e.i.Eval(dirName + ".Run")
	if err != nil {
		e.l.Error().Err(err).Msgf("Failed to get %s.Run() function!", dirName)
		return PluginHandle{}
	}

	pluginHandle := PluginHandle{
		PackageName: dirName,
		Run:         v.Interface().(func([]byte) (msg []byte, err error)),
	}
	return pluginHandle
}

// 加载所有插件文件
// pluginsDir: 插件目录，遍历出所有插件目录，然后调用 Load() 加载单个插件文件
// 返回所有插件的入口函数句柄列表
func (e *Engine) LoadAll(pluginsDir string) []PluginHandle {
	plugins := []PluginHandle{}

	dirs, err := tool.GetSubDirectories(pluginsDir)
	if err != nil {
		e.l.Error().Err(err).Msgf("Failed to glob %s", filepath.Join(pluginsDir, "*"))
		return nil
	}

	for _, dir := range dirs {
		e.l.Info().Msgf("Loading plugin %s", dir)
		h := e.Load(dir)
		if h.Run != nil {
			plugins = append(plugins, h)
		}
	}

	e.l.Info().Msgf("Loaded %d plugins", len(plugins))

	return plugins
}
