package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ghc-golang-hoangth7/finalprj/client/graph"
	"github.com/ghc-golang-hoangth7/finalprj/common"
	pbFlights "github.com/ghc-golang-hoangth7/finalprj/pb/flights"
	pbPlanes "github.com/ghc-golang-hoangth7/finalprj/pb/planes"
)

func main() {
	appConfig, err := common.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	planesConn, err := grpc.Dial(appConfig.GetPlanesAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	flightsConn, err := grpc.Dial(appConfig.GetFlightsAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer planesConn.Close()
	defer flightsConn.Close()

	planesServiceClient := pbPlanes.NewPlanesServiceClient(planesConn)
	flightsServiceClient := pbFlights.NewFlightServiceClient(flightsConn)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			FlightsService: flightsServiceClient,
			PlanesService:  planesServiceClient,
		},
		Directives: graph.DirectiveRoot{
			Validate: graph.Validate,
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://%v:%v/ for GraphQL playground", appConfig.GraphQLHost, appConfig.GraphQLPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", appConfig.GraphQLHost, appConfig.GraphQLPort), nil))
}
