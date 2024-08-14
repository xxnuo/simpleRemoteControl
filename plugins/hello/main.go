package hello

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Run(jsonData []byte) (msg []byte, err error) {
	type helloData struct {
		Name string `json:"name"`
	}
	var data helloData

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	if data.Name == "" {
		return nil, errors.New("name is required")
	}

	fmt.Printf("Hello, %s!", data.Name)
	return []byte("Got it!"), nil
}
