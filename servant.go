package main

import (
	"flag"
	"fmt"
	. "github.com/jserver/serverstyle/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

var (
	host = flag.String("host", "localhost", "Host to listen on")
	port = flag.Int("port", 1234, "Port to bind to")
)

func main() {
	flag.Parse()
	address := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Println("Address: " + address)

	err := os.Mkdir(os.Getenv("HOME") + "/style_scripts", 0755)
	if err != nil {
		fmt.Println("Unable to create script dir")
	}

	install := new(AptGetInstall)
	rpc.Register(install)

	update := new(AptGetUpdate)
	rpc.Register(update)

	upgrade := new(AptGetUpgrade)
	rpc.Register(upgrade)

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
