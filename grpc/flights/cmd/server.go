package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/ghc-golang-hoangth7/finalprj/grpc/flights/handlers"
	pb "github.com/ghc-golang-hoangth7/finalprj/pb/flights"
	pbPlanes "github.com/ghc-golang-hoangth7/finalprj/pb/planes"
)

// Db connection info
// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "flights_db"
// )

func main() {
	// Connect to the database
	connStr := "postgres://postgres:postgres@localhost:5432/planes_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	planesConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer planesConn.Close()

	// Create a new flight service instance
	flightService := handlers.NewFlightService(db, pbPlanes.NewPlanesServiceClient(planesConn))

	// Start a gRPC server on port 50051
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFlightServiceServer(grpcServer, flightService)
	log.Println("Server started at port :50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
