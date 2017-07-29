package go_serial

import (
	"io/ioutil"
	"encoding/json"
    "os"
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
    raw, err := ReadConfigFile("config.json")
    if nil != err {
        panic(err)
    }
    config := new(Config)
    json.Unmarshal(raw, &config)
    return config
}

func ReadConfigFile(filename string) ([]byte, error) {
    if false == ConfigExists(filename) {
        return nil, nil
    }
    return ioutil.ReadFile(filename)
}

func ConfigExists(filename string) (bool) {
    stat, err := os.Stat(filename)
    if nil != err {
        return false
    }
    return stat.Mode().IsRegular()
}