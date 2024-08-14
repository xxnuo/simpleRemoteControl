//go:build ignore

package walk

import (
	"fmt"
	"os"
	"path/filepath"
)

func Run(jsonData string) (msg string, err error) {
	root := "./" // 你可以将这里替换为你要遍历的目录路径

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fmt.Print(path)
		return err
	})

	if err != nil {
		fmt.Printf("遍历目录时发生错误: %v\n", err)
	}

	return "Got it!", nil
}
