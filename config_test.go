package go_serial

import (
    "testing"
    "os"
)

func TestWriteConfig(t *testing.T) {
    beforeTest()
    config := Config{Min: 1}
	WriteConfig("", config)
}

func TestConfigExists(t *testing.T) {
    beforeTest()
    os.Create("config.json")
    exists := ConfigExists("config.json")
    if true != exists {
        t.Error("File should exist")
    }
}

func TestConfigExistsOnMissingFile(t *testing.T) {
    exists := ConfigExists("/dev/false")
    if true == exists {
        t.Error("File should not exist")
    }
}

func TestReadConfigFileReturnsNoError(t *testing.T) {
    _, err := ReadConfigFile("/dev/false")
    if nil != err {
      t.Error(err, "!=", os.ErrNotExist)
    }
}

func TestReadConfigFileReturnsNil(t *testing.T) {
    raw, _ := ReadConfigFile("/dev/false")
    if nil != raw {
        t.Error("raw is not nil when file does not exist")
    }
}

func TestReadConfig(t *testing.T) {
    beforeTest()
    config := ReadConfig("",)
    if config.Min != 0 {
        t.Error("Min is not valid")
    }
}

func TestReadConfigAfterWrite(t *testing.T) {
    beforeTest()
    configOriginal := Config{Min: 2017}
    WriteConfig("", configOriginal)
    configResponse := ReadConfig("")
    if configResponse.Min != configOriginal.Min {
        t.Error("Min is not valid")
    }
}

func beforeTest() {
    os.Remove("config.json")
}