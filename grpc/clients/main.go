package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ghc-golang-hoangth7/finalprj/pb/flights"
	"github.com/ghc-golang-hoangth7/finalprj/pb/planes"
	"github.com/ghc-golang-hoangth7/finalprj/utils"
)

func main() {
	// testPlanes()
	testFlights()
}

func testFlights() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a new plane management client
	client := flights.NewFlightServiceClient(conn)

	flight := flights.Flight{
		PlaneNumber:          "DEF123",
		DeparturePoint:       "A",
		DestinationPoint:     "B",
		DepartureTime:        timestamppb.New(time.Now().AddDate(0, 0, 7)),
		EstimatedArrivalTime: timestamppb.New(time.Now().AddDate(0, 0, 7).Add(4 * time.Hour)),
		AvailableSeats:       400,
	}

	if id, err := client.UpsertFlight(context.Background(), &flight); err != nil {
		log.Fatal(err.Error())
	} else {
		flight.Id = id.Id
		log.Printf("Flight created: \n%v", utils.ObjToString(flight.ToModels()))
	}
	if _, err := client.BookFlight(context.Background(), &flights.BookFlightRequest{
		FlightId:   flight.Id,
		SeatNumber: 4,
	}); err != nil {
		log.Fatal(err.Error())
	}
}

func testPlanes() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a new plane management client
	client := planes.NewPlanesServiceClient(conn)

	// Add or update a plane
	planeId := uuid.New().String()
	plane := &planes.Plane{
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

func addOrUpdatePlane(client planes.PlanesServiceClient, plane *planes.Plane) {
	// Add or update the plane
	_, err := client.UpsertPlane(context.Background(), plane)
	if err != nil {
		log.Fatalf("Failed to add or update plane: %v", err)
	}
	fmt.Printf("Added or updated plane with ID: %s\n", plane.PlaneId)
}

func listPlanes(client planes.PlanesServiceClient) {
	// List planes
	fmt.Println("Viewing a list of planes...")
	planes, err := client.GetPlanesList(context.Background(), &planes.Plane{
		PlaneNumber: "DEF123",
	})
	if err != nil {
		log.Fatalf("Failed to list planes: %v", err)
	}
	for _, plane := range planes.Planes {
		fmt.Printf("%s, %s, %d, %s\n", plane.PlaneId, plane.PlaneNumber, plane.TotalSeats, plane.Status)
	}
}

func updatePlaneStatus(client planes.PlanesServiceClient, planeId string, status string) {
	_, err := client.ChangePlaneStatus(context.Background(), &planes.PlaneStatusRequest{
		PlaneId: planeId,
		Status:  status,
	})
	if err != nil {
		log.Fatalf("Failed to update plane status: %v", err)
	}
	fmt.Printf("Updated status of plane with ID %s to %s\n", planeId, status)
}
