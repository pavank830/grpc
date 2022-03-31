package main

import (
	"flag"
	"log"

	"github.com/grpc/server"
)

func main() {
	var portFlag = flag.Int("port", 8080, "default port 8080")
	flag.Parse()
	log.Printf("server port:%d", *portFlag)
	server.StartGRPCSvr(*portFlag)
}
