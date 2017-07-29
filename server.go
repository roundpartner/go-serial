package go_serial

import (
    "net/http"
    "fmt"
    "encoding/json"
)

var config = &Config{}

var configWorkingDirectory = ""

func StartServer(workingDirectory string) *http.Server {
    configWorkingDirectory = workingDirectory
    config = ReadConfig(workingDirectory)

    http.HandleFunc("/", handler)
    server := &http.Server{Addr: "127.0.0.1:53231"}
    server.ListenAndServe()
    return server
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)

    config.Min += 1
    configJson, _ := json.Marshal(config)

    fmt.Fprint(w, string(configJson))

    WriteConfig(configWorkingDirectory, *config)
}