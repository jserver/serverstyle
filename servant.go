package main

import (
	"flag"
	"fmt"
	. "github.com/jserver/serverstyle/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

var (
	host = flag.String("host", "localhost", "Host to listen on")
	port = flag.Int("port", 1234, "Port to bind to")
)

func main() {
	flag.Parse()
	address := fmt.Sprintf("%s:%d", *host, *port) 
	fmt.Println("Address: " + address)

	aptGet := new(AptGet)
	rpc.Register(aptGet)

	script := new(Script)
	rpc.Register(script)

	test := new(Test)
	rpc.Register(test)

	rpc.HandleHTTP()
	l, e := net.Listen("tcp", address)
	if e != nil {
		log.Fatal("listen error:", e)
	}

	fmt.Println("Server Started...")
	http.Serve(l, nil)
}
