package main

import (
	"grepRPC"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Missing port number: ./grepServer <Port Num>")
	}
	port := os.Args[1]
	grepCmd := new(grepRPC.GrepRes)

	rpc.Register(grepCmd)
	rpc.HandleHTTP()
	// server start listening
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("listen error: ", err)
	}
	// serve client request
	for {
		http.Serve(listen, nil)
	}
}
