package go_serial

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Min  int `json:"min"`
}

func WriteConfig() {
	config := Config{Min: 1}
	configJson, _ := json.Marshal(config)
	ioutil.WriteFile("output.json", configJson, 0644)
}
