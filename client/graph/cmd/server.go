package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"google.golang.org/grpc"

	"github.com/ghc-golang-hoangth7/finalprj/client/graph"
	pbFlights "github.com/ghc-golang-hoangth7/finalprj/pb/flights"
	pbPlanes "github.com/ghc-golang-hoangth7/finalprj/pb/planes"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	planesConn, err := grpc.Dial("planes-service-address:port", grpc.WithInsecure())
	flightsConn, err := grpc.Dial("flights-service-address:port", grpc.WithInsecure())
	defer planesConn.Close()
	defer flightsConn.Close()

	if err != nil {
		// Handle connection error
	}
	planesServiceClient := pbPlanes.NewPlanesServiceClient(planesConn)

	if err != nil {
		// Handle connection error
	}
	flightsServiceClient := pbFlights.NewFlightServiceClient(flightsConn)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		FlightsService: flightsServiceClient,
		PlanesService:  planesServiceClient,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
