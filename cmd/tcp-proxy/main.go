package main

import (
	"runtime"

	"tcp-proxy/internal/cli"
	"tcp-proxy/internal/proxy"
)

func main() {
	runtime.GOMAXPROCS(1)
	port := cli.Parse()
	proxy.Start(port)
}
