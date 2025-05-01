package user

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kwul0208/gateway/pkg/user/pb"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	fmt.Println("API Gateway : InitAuthMiddleware")
	return AuthMiddlewareConfig{svc: svc}
}

func (c *AuthMiddlewareConfig) UserAuth(ctx *gin.Context) {
	fmt.Println("API Gateway: AuthRequired")
	authorization := ctx.Request.Header.Get("Authorization")

	if authorization == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, no Authorization header provided",
		})
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) > 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, invalid Authorization header",
		})
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
		Role:  "user",
	})

	log.Printf("res")
	log.Printf(res.String())

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, invalid token",
		})
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()

}
