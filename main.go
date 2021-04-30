package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/anchamber/genetics-system/db"
	pb "github.com/anchamber/genetics-system/proto"
	"github.com/anchamber/genetics-system/service"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	enableTls       = flag.Bool("enable_tls", false, "Use TLS - required for HTTP2.")
	tlsCertFilePath = flag.String("tls_cert_file", "./crt/localhost.crt", "Path to the CRT/PEM file.")
	tlsKeyFilePath  = flag.String("tls_key_file", "./crt/localhost.key", "Path to the private key file.")
	gRPCPort        = flag.Int("grpc-port", 10000, "The gRPC server port")
	gRPCWebPort     = flag.Int("grpc-web-port", 10001, "The gRPC web server port")
)

func main() {
	log.Printf("gRPC port: %d | gRPC web port: %d\n", *gRPCPort, *gRPCWebPort)
	addr := fmt.Sprintf(":%d", *gRPCPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	s := grpc.NewServer()
	pb.RegisterSystemServiceServer(s, service.New(db.NewMockDB(nil)))

	// Serve gRPC Server
	// go func() {
	log.Printf("Starting gRPC server %s\n", addr)
	log.Fatal(s.Serve(lis))
	// }()
	// startGRPCWeb()
	//wrappedServer := grpcweb.WrapServer(s)
	//handler := func(resp http.ResponseWriter, req *http.Request) {
	//	wrappedServer.ServeHTTP(resp, req)
	//}
	//webAddr := fmt.Sprintf(":%d", *gRPCWebPort)
	//srv := &http.Server{
	//	Handler: http.HandlerFunc(handler),
	//	Addr:    webAddr,
	//}
	//// Serve the webapp over TLS
	//log.Printf("Starting gRPC web server %s\n", webAddr)
	//log.Fatal(srv.ListenAndServe())
}

func startGRPCWeb() {

	flag.Parse()

	port := *gRPCWebPort

	grpcServer := grpc.NewServer()
	pb.RegisterSystemServiceServer(grpcServer, service.New(db.NewMockDB(nil)))
	grpclog.SetLogger(log.New(os.Stdout, "exampleserver: ", log.LstdFlags))

	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}

	grpclog.Printf("Starting server. http port: %d, tls enabled %v", port, *enableTls)

	if *enableTls {
		if err := httpServer.ListenAndServeTLS(*tlsCertFilePath, *tlsKeyFilePath); err != nil {
			grpclog.Fatalf("failed starting http2 server: %v", err)
		}
	} else {
		if err := httpServer.ListenAndServe(); err != nil {
			grpclog.Fatalf("failed starting http server: %v", err)
		}
	}

}
