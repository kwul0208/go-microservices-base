package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kwul0208/gateway/pkg/user"
	"github.com/kwul0208/gateway/pkg/user/config"
)

func main() {
	log.Print("Starting API Gateway")

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Failed load config", err)
	}

	r := gin.Default()

	userSvc := *user.RegisterRoutes(r, &c)
	log.Println(userSvc)
	r.Run()
}
