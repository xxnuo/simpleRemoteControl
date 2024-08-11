package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xxnuo/simpleRemoteControl/internal/log"
)

// cobra.go
type CobraConfig struct {
	Addr      string
	Port      int
	File      string
	IsJsonLog bool
}

// global variables
var (
	Cfg       CobraConfig // cobra.go
	Logger    log.Logger  // internal/log/log.go
	ApiServer *fiber.App
)
