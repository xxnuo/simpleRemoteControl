package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func main() {
	i := interp.New(interp.Options{})
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

	v, err := i.Eval("plugin.Init")
	if err != nil {
		panic(err)
	}

	Init := v.Interface().(func())
	Init()
	// r := Init()
	// println(r)
}
