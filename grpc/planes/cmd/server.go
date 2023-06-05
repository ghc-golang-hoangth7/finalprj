package main

import (
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/ghc-golang-hoangth7/finalprj/common"
	"github.com/ghc-golang-hoangth7/finalprj/grpc/planes/handlers"
	pbFlights "github.com/ghc-golang-hoangth7/finalprj/pb/flights"
	pb "github.com/ghc-golang-hoangth7/finalprj/pb/planes"
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

	flightSrv, flightSrvConn, err := pbFlights.NewServiceClient(appConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer flightSrvConn.Close()

	// Create a new plane service instance
	planeService := handlers.NewPlanesService(db, flightSrv)

	// Start a gRPC server on port 50051
	listener, err := net.Listen("tcp", appConfig.GetPlanesAddr())
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPlanesServiceServer(grpcServer, planeService)
	log.Printf("Server started at port :%v", appConfig.PlanesPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
