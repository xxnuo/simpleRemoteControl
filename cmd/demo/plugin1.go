//go:build ignore

package plugin1

import (
	"runtime"
)

func Init() {
	print("Calling ext1" + " running on " + runtime.GOOS + runtime.GOARCH)
}
func main() {
	print("main from ext1")
}
func init() {
	print("init from ext1")
}
