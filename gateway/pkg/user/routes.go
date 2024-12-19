package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kwul0208/gateway/pkg/config"
	"github.com/kwul0208/gateway/pkg/user/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	log.Print("API Gateway: Register Routes")

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	user := r.Group("/auth")
	user.POST("/register", svc.Register)
	user.POST("/login", svc.Login)

	return svc

}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}
func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
