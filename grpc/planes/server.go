package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/ghc-golang-hoangth7/finalprj/grpc/planes/handlers"
	pb "github.com/ghc-golang-hoangth7/finalprj/pb/planes"
)

// Db connection info
// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "planes_db"
// )

func main() {
	// Connect to the database
	connStr := "postgres://postgres:postgres@localhost:5432/planes_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create a new planeService instance
	planeService := handlers.NewPlanesService(db)

	// Start a gRPC server on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPlanesServiceServer(grpcServer, planeService)
	log.Println("Server started at port :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
