package main

import (
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/ghc-golang-hoangth7/finalprj/common"
	"github.com/ghc-golang-hoangth7/finalprj/grpc/flights/handlers"
	pb "github.com/ghc-golang-hoangth7/finalprj/pb/flights"
	pbPlanes "github.com/ghc-golang-hoangth7/finalprj/pb/planes"
)

func main() {
	appConfig, err := common.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := common.InitDb(appConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	planesSrv, planesConn, err := pbPlanes.NewServiceClient(appConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer planesConn.Close()

	// Create a new flight service instance
	flightService := handlers.NewFlightService(db, planesSrv)

	// Start a gRPC server on port 50051
	listener, err := net.Listen("tcp", appConfig.GetFlightsAddr())
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFlightServiceServer(grpcServer, flightService)
	log.Printf("Server started at port :%v", appConfig.FlightsPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
