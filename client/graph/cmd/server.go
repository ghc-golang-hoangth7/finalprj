package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

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
	planesServiceClient, planesConn, err := pbPlanes.NewServiceClient(appConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer planesConn.Close()

	flightsServiceClient, flightsConn, err := pbFlights.NewServiceClient(appConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer flightsConn.Close()

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
