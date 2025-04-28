package main

import (
	"fmt"
	"log"
	"net"
	"os"

	csipb "github.com/Shravani-Phadol/csi-driver/api"
	"github.com/Shravani-Phadol/csi-driver/driver"
	"google.golang.org/grpc"
)

func main() {
	// Define server socket location
	socket := "/tmp/csi.sock"
	// Remove any existing socket
	os.Remove(socket)
	// Create a listener for the socket
	listener, err := net.Listen("unix", socket)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Create a gRPC server
	grpcServer := grpc.NewServer()
	// Register the identity, controller, and node servers
	csipb.RegisterIdentityServer(grpcServer, &driver.IdentityServer{})
	csipb.RegisterControllerServer(grpcServer, &driver.ControllerServer{})
	csipb.RegisterNodeServer(grpcServer, &driver.NodeServer{})
	// Start the gRPC server
	fmt.Printf("CSI driver is listening on %s\n", socket)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
