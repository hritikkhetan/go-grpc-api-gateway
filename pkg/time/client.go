package time

import (
	"fmt"

	"github.com/hritikkhetan/go-grpc-api-gateway/pkg/config"
	"github.com/hritikkhetan/go-grpc-api-gateway/pkg/time/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.TimeServiceClient
}

func InitServiceClient(c *config.Config) pb.TimeServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.TimeSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewTimeServiceClient(cc)
}
