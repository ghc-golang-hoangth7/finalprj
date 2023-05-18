package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ghc-golang-hoangth7/finalprj/pb/planes"
	"github.com/google/uuid"
)

func main() {
	testPlanes()
}

func testPlanes() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a new plane management client
	client := pb.NewPlanesServiceClient(conn)

	// Add or update a plane
	planeId := uuid.New().String()
	plane := &pb.Plane{
		PlaneId:     planeId,
		PlaneNumber: "DEF123",
		TotalSeats:  240,
		Status:      "ready",
	}
	addOrUpdatePlane(client, plane)

	// List planes
	listPlanes(client)

	// Update the status of a plane
	updatePlaneStatus(client, planeId, "flying")
}

func addOrUpdatePlane(client pb.PlanesServiceClient, plane *pb.Plane) {
	// Add or update the plane
	_, err := client.AddOrUpdatePlane(context.Background(), plane)
	if err != nil {
		log.Fatalf("Failed to add or update plane: %v", err)
	}
	fmt.Printf("Added or updated plane with ID: %s\n", plane.PlaneId)
}

func listPlanes(client pb.PlanesServiceClient) {
	// List planes
	fmt.Println("Viewing a list of planes...")
	planes, err := client.ListPlanes(context.Background(), &pb.Plane{
		PlaneNumber: "DEF123",
	})
	if err != nil {
		log.Fatalf("Failed to list planes: %v", err)
	}
	for _, plane := range planes.Planes {
		fmt.Printf("%s, %s, %d, %s\n", plane.PlaneId, plane.PlaneNumber, plane.TotalSeats, plane.Status)
	}
}

func updatePlaneStatus(client pb.PlanesServiceClient, planeId string, status string) {
	_, err := client.UpdatePlaneStatus(context.Background(), &pb.Plane{
		PlaneId: planeId,
		Status:  status,
	})
	if err != nil {
		log.Fatalf("Failed to update plane status: %v", err)
	}
	fmt.Printf("Updated status of plane with ID %s to %s\n", planeId, status)
}
