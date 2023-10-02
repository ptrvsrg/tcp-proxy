package main

import (
	"runtime"

	proxy "tcp-proxy/internal"
	cli "tcp-proxy/internal/cli-parser"
)

func main() {
	runtime.GOMAXPROCS(1)
	port := cli.Parse()
	proxy.Start(port)
}
