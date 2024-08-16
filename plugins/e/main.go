package e

import (
	"errors"
)

func Run(data string) (msg string, err error) {
	err = errors.New("error test")
	return err.Error(), err
}
