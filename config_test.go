package go_serial

import (
    "testing"
    "os"
)

func TestWriteConfig(t *testing.T) {
	WriteConfig()
}

func TestReadConfig(t *testing.T) {
    ReadConfig()
}

func TestConfigExists(t *testing.T) {
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