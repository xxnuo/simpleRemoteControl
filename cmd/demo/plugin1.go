//go:build ignore

package plugin

import (
	"runtime"
)

func Init() {
	print("Init from ext1" + " running on " + runtime.GOOS + runtime.GOARCH)
}
func Main() {
	print("Main from ext1")
}
func End() {
	print("End from ext1")
}
