package main

import (
	"flag"
	"log"

	"github.com/grpc/client"
)

var (
	userIDFlag     *int
	userIDListFlag *string
	serverFlag     *string
)

func init() {
	userIDFlag = flag.Int("id", 0, "User ID , --id 10")
	userIDListFlag = flag.String("idlist", "", "User ID List , --idlist 1,2,3,4")
	serverFlag = flag.String("server", "", "grpc server addr")
	flag.Parse()
	if *serverFlag == "" {
		log.Fatalf("grpc server addr is complusory")
	}
	log.Printf("grpc server addr:%s", *serverFlag)
	if *userIDFlag != 0 {
		log.Printf("user id:%d", *userIDFlag)
	}
	if *userIDListFlag != "" {
		log.Printf("user id list:%s", *userIDListFlag)
	}

}
func main() {
	client.Trigger(*userIDFlag, *userIDListFlag, *serverFlag)
}
