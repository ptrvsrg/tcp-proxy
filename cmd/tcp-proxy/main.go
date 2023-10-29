package main

import (
	"runtime"

	"tcp-proxy/internal/cli"
	"tcp-proxy/internal/proxy"
)

const title string = "  ________________                                \n" +
	" /_  __/ ____/ __ \\   ____  _________  _  ____  __\n" +
	"  / / / /   / /_/ /  / __ \\/ ___/ __ \\| |/_/ / / /\n" +
	" / / / /___/ ____/  / /_/ / /  / /_/ />  </ /_/ / \n" +
	"/_/  \\____/_/      / .___/_/   \\____/_/|_|\\__, /  \n" +
	"                  /_/                    /____/   \n" +
	"tcp-proxy:1.0.0                                   \n"

func main() {
	println(title)
	runtime.GOMAXPROCS(1)
	port := cli.Parse()
	proxy.Start(port)
}
