package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/anchamber/genetics-system/proto"
	"github.com/anchamber/genetics-system/server"
	"google.golang.org/grpc"
)

var (
	gRPCPort    = flag.Int("grpc-port", 10000, "The gRPC server port")
	gatewayPort = flag.Int("gateway-port", 11000, "The gRPC-Gateway server port")
)

func main() {
	addr := fmt.Sprintf("localhost:%d", *gRPCPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	s := grpc.NewServer()
	pb.RegisterSystemServiceServer(s, server.New())

	// Serve gRPC Server
	fmt.Printf("Serving gRPC on https://%s\n", addr)
	// go func() {
	// 	log.Fatal(s.Serve(lis))
	// }()
	log.Fatal(s.Serve(lis))
}
