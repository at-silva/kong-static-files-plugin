package main

import (
	"github.com/Kong/go-pdk/server"
	"github.com/at-silva/kong-plugin-static-files/plugin"
)

const (
	version  = "0.1"
	priority = 1000
)

func main() {
	if err := server.StartServer(plugin.New, version, priority); err != nil {
		panic(err)
	}
}
