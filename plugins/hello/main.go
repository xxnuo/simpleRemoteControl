package hello

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Run(jsonData string) (msg string, err error) {
	type helloData struct {
		Name string `json:"name"`
	}
	var data helloData

	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return jsonData, err
	}

	if data.Name == "" {
		return jsonData, errors.New("name is required")
	}

	fmt.Printf("Hello, %s!", data.Name)
	return "Hi, Dear " + data.Name + "!", nil
}
