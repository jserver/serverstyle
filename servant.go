package main

import (
	"flag"
	"fmt"
	. "github.com/jserver/serverstyle/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
	//"os"
	"strconv"
)

var (
	host = flag.String("host", "localhost", "Host to listen on")
	port = flag.Int("port", 1234, "Port to bind to")
)

func main() {
	flag.Parse()
	address := *host + ":" + strconv.Itoa(*port) 
	fmt.Println("Address: " + address)

	/*
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal("addresses fail")
	}
	for _, addr := range addresses {
		fmt.Println("ADDRESS:" + addr.String())
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("hostname fail")
	}
	fmt.Println("HOST:" + hostname)
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		log.Fatal("lookup fail")
	}
	for _, addr := range addrs {
		fmt.Println("ADDR:" + addr)
	}
	*/

	aptGet := new(AptGet)
	rpc.Register(aptGet)

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
