package server

import (
	"fmt"
	"log"
	"net"

	"github.com/grpc/user"
	"google.golang.org/grpc"
)

//StartGRPCSvr -- start grpc server
func StartGRPCSvr(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := user.Server{}
	grpcServer := grpc.NewServer()
	user.RegisterUserDetailsServer(grpcServer, &s)
	log.Printf("Starting to serve on %s\n", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
