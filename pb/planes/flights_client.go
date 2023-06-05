package planes

import (
	"github.com/ghc-golang-hoangth7/finalprj/common"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServiceClient(config *common.Config) (PlanesServiceClient, *grpc.ClientConn, error) {
	// Set up a gRPC connection to the Planes service
	conn, err := grpc.Dial(config.GetPlanesAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	// Create the PlanesServiceClient using the gRPC connection
	client := NewPlanesServiceClient(conn)

	return client, conn, nil
}
