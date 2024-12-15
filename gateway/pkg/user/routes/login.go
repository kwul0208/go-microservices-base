package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwul0208/gateway/pkg/user/pb"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.UserServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.Bind(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
