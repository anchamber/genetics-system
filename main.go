package main

import (
	"fmt"
	"github.com/anchamber/genetics-system/db"
	pb "github.com/anchamber/genetics-system/proto"
	"github.com/anchamber/genetics-system/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	configuration := LoadConfiguration()

	addr := fmt.Sprintf(":%s", configuration.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	s := grpc.NewServer()
	pb.RegisterSystemServiceServer(s, service.NewSystemService(db.NewSystemDBMock(nil)))
	pb.RegisterTankServiceServer(s, service.NewTankService(db.NewTankDBMock(nil)))

	// Serve gRPC Server
	log.Printf("Starting gRPC server %s\n", addr)
	log.Fatal(s.Serve(lis))
}
