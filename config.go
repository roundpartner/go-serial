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
	ioutil.WriteFile("config.json", configJson, 0644)
}

func ReadConfig() *Config {
    raw, err := ioutil.ReadFile("config.json")
    if nil != err {
        panic(err)
    }
    config := new(Config)
    json.Unmarshal(raw, &config)
    return config
}
