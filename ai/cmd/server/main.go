package main

import (
	"flag"

	"github.com/tmadeira/hex/ai/server"
)

var (
	port = flag.Int("port", 8080, "tcp port to listen")
)

func main() {
	flag.Parse()
	server.Run(*port)
}
