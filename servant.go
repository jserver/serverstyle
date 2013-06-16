package main

import (
	"flag"
	"fmt"
	. "github.com/jserver/serverstyle/server"
)

var (
	host = flag.String("host", "localhost", "Host to listen on")
	port = flag.Int("port", 1234, "Port to bind to")
)

func main() {
	flag.Parse()
	address := fmt.Sprintf("%s:%d", *host, *port)

	StartServer(address)
}
