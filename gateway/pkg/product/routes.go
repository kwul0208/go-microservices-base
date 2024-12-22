package product

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kwul0208/gateway/pkg/config"
	"github.com/kwul0208/gateway/pkg/product/routes"
	"github.com/kwul0208/gateway/pkg/user"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, userSvc *user.ServiceClient) {
	a := user.InitAuthMiddleware(userSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	route := r.Group("/product")

	route.GET("/:id", svc.FindOne)
	route.GET("/", svc.FindAll)

	route.Use(a.UserAuth)
	route.POST("/", svc.CreateProduct)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	fmt.Println("API Gateway : FindOne")
	routes.FindOne(ctx, svc.Client)
}
func (svc *ServiceClient) FindAll(ctx *gin.Context) {
	fmt.Println("API Gateway : FindOne")
	routes.FindAll(ctx, svc.Client)
}
func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
