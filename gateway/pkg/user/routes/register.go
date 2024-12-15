package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwul0208/gateway/pkg/user/pb"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, c pb.UserServiceClient) {
	b := RegisterRequestBody{}

	if err := ctx.Bind(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
