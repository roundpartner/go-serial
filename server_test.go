package go_serial

import (
    "testing"
    "net/http"
    "strings"
    "io/ioutil"
    "encoding/json"
    "time"
    "os"
    "io"
)

func TestStartServerStarts(t *testing.T) {
    os.Remove("config.json")
    go StartServer()
    time.Sleep(10 * time.Millisecond)
    _, err := http.Get("http://127.0.0.1:53231")
    if err != nil {
        t.Error(err)
    }
}

func TestStartServerRespondsOK(t *testing.T) {
    resp, _ := http.Get("http://127.0.0.1:53231")
    if resp.Status != "200 OK" {
        t.Error("Response was not 200 OK got", resp.Status, "instead")
    }
}

func TestStartServerRespondsJson(t *testing.T) {
    resp, _ := http.Get("http://127.0.0.1:53231")
    if false == strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
        t.Error("Response did not contain application/json got", resp.Header.Get("Content-Type"), "instead")
    }
}

func TestStartServerResponseParses(t *testing.T) {
    config := getJsonResponse()
    if config.Min == 0 {
        t.Error(config.Min)
    }
}

func TestStartServerResponseCounts(t *testing.T) {
    config := getJsonResponse()
    configNext := getJsonResponse()
    if config.Min >= configNext.Min {
        t.Error("config was not ", config.Min, " < ", configNext.Min)
    }
}

func getJsonResponse() *Config {
    resp, _ := http.Get("http://127.0.0.1:53231")
    body, _ := ioutil.ReadAll(resp.Body)
    config := new(Config)
    json.Unmarshal([]byte(body), &config)
    return config
}
