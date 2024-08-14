package tool

import (
	"os"
	"path/filepath"
)

func GetSubDirectories(dir string) ([]string, error) {
	var subDirs []string

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && path != dir {
			subDirs = append(subDirs, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return subDirs, nil
}
