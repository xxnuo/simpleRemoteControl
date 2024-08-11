// #!/usr/bin/env yaegi

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "../" // 你可以将这里替换为你要遍历的目录路径

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fmt.Println(path)
		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录时发生错误: %v\n", err)
	}
}
