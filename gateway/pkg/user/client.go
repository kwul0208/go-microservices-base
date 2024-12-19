package user

import (
	"log"

	"github.com/kwul0208/gateway/pkg/config"
	"github.com/kwul0208/gateway/pkg/user/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.UserServiceClient
}

func InitServiceClient(c *config.Config) pb.UserServiceClient {
	log.Print("API Gateway: Init Service Client")
	//	using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthSuvUrl, grpc.WithInsecure())

	if err != nil {
		log.Panicln("Could not connect:", err)
	}
	return pb.NewUserServiceClient(cc)

}
