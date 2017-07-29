package go_serial

import (
    "net/http"
    "fmt"
    "encoding/json"
)

var config = Config{Min: 1}

func StartServer() *http.Server {
    http.HandleFunc("/", handler)
    server := &http.Server{Addr: "127.0.0.1:53231"}
    go server.ListenAndServe()
    return server

}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)

    config.Min += 1
    configJson, _ := json.Marshal(config)

    fmt.Fprint(w, string(configJson))
}