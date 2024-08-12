//go:build ignore

package plugin

import (
	"runtime"

	"github.com/rs/zerolog/log"
)

func Init() {
	println("Init from ext1" + " running on " + runtime.GOOS)
	log.Info("hi")
}
func Main() { println("Main from ext1") }
func End()  { println("End from ext1") }
