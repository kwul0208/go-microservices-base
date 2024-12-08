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
	userServiceAddr    = "localhost:2001"
)

func main() {
	// -- GRPC CONNECTION INITIAL --
	ProductConn, err := grpc.Dial(productServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer ProductConn.Close()

	log.Println("Dialing product service at ", productServiceAddr)
	productClient := pb.NewProductServiceClient(ProductConn)

	UserConn, err := grpc.Dial(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer UserConn.Close()

	log.Println("Dialing product service at ", userServiceAddr)
	userClient := pb.NewUserServiceClient(UserConn)
	// -- END --

	// Use Gin as the router
	r := gin.Default()

	// Initialize handler and register routes with Gin
	productHandler := handler.NewHandler(productClient)
	productRoute := r.Group("/v1/api/product")
	{
		productRoute.GET("/", productHandler.HandlerGetProduct)
		productRoute.GET("/detail", productHandler.HandlerGetProductById)
		productRoute.POST("/store", productHandler.HandleCreateProduct)
		productRoute.PUT("/update", productHandler.HandleUpdateProduct)
		productRoute.DELETE("delete", productHandler.HandleDeleteProduct)
	}

	userHandler := handler.NewUserHandler(userClient)
	userRoute := r.Group("/v1/api/user")
	{
		userRoute.POST("/register", userHandler.HandleRegister)
		userRoute.POST("/login", userHandler.HandlerLogin)
	}

	log.Printf("Starting server at %s", httpAddr)
	if err := r.Run(httpAddr); err != nil {
		log.Fatal("Failed to start the server")
	}
}
