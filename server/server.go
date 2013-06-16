package server

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type Results interface {
	GetErr() string
	GetOutput() string
}

var	logger *log.Logger

func StartServer(address string) {

	logfile, err := os.OpenFile("/var/log/serverstyle.log", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	logger = log.New(logfile, "", log.LstdFlags)

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
		logger.Fatal("listen error:", e)
	}

	logger.Println("Server Started...")
	http.Serve(l, nil)
}
