package main

import (
    "github.com/artyom/autoflags"
    "flag"
    "github.com/roundpartner/go-serial"
)

var autoFlagsConfig = struct {
    WorkingDirectory    string `flag:"wd,working directory for config file"`
}{
    WorkingDirectory: "",
}

func main() {
    autoflags.Define(&autoFlagsConfig)
    flag.Parse()
    go_serial.StartServer(autoFlagsConfig.WorkingDirectory)
}
