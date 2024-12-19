package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwul0208/gateway/pkg/product/pb"
)

func FindAll(ctx *gin.Context, c pb.ProductServiceClient) {
	res, err := c.FindAll(context.Background(), &pb.FindAllRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
