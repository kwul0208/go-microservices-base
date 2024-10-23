package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kwul0208/common"
	pb "github.com/kwul0208/common/api"
	"github.com/kwul0208/gateway/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr           = common.EnvString("HTTP_ADDR", ":3000")
	productServiceAddr = "localhost:2000"
)

func main() {
	conn, err := grpc.Dial(productServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	log.Println("Dialing product service at ", productServiceAddr)
	c := pb.NewProductServiceClient(conn)

	// Use Gin as the router
	r := gin.Default()

	// Initialize handler and register routes with Gin
	handlerInit := handler.NewHandler(c)
	productRoute := r.Group("/v1/api/product")
	{
		productRoute.GET("/", handlerInit.HandlerGetProduct)
		productRoute.GET("/detail", handlerInit.HandlerGetProductById)
		productRoute.POST("/store", handlerInit.HandleCreateProduct)
		productRoute.PUT("/update", handlerInit.HandleUpdateProduct)
		productRoute.DELETE("delete", handlerInit.HandleDeleteProduct)
	}

	log.Printf("Starting server at %s", httpAddr)
	if err := r.Run(httpAddr); err != nil {
		log.Fatal("Failed to start the server")
	}
}
