package flights

import (
	"github.com/ghc-golang-hoangth7/finalprj/common"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServiceClient(config common.Config) (FlightServiceClient, *grpc.ClientConn, error) {
	// Set up a gRPC connection to the Flights service
	conn, err := grpc.Dial(config.GetFlightsAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	// Create the FlightsServiceClient using the gRPC connection
	client := NewFlightServiceClient(conn)

	return client, conn, nil
}
