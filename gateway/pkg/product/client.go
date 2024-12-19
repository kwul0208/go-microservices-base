package product

import (
	"fmt"

	"github.com/kwul0208/gateway/pkg/config"
	"github.com/kwul0208/gateway/pkg/product/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	fmt.Println("API Gateway : InitServiceClient")

	cc, err := grpc.Dial(c.ProductSuvUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect: ", err)
	}

	return pb.NewProductServiceClient(cc)
}
