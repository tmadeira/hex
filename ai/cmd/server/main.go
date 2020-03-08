package main

import (
	"flag"

	"github.com/tmadeira/hex/ai/server"
)

var (
	port      = flag.Int("port", 8080, "tcp port to listen")
	playerID  = flag.Int("pid", 2, "ai player id (1 or 2)")
	strategy  = flag.String("strategy", "minimax", "ai strategy")
	heuristic = flag.String("heuristic", "mindistance", "minimax heuristic to use")
)

func main() {
	flag.Parse()
	server.Run(*port, *playerID, *strategy, *heuristic)
}
