package main

import (
	"context"
	"log"
	"net"

	"github.com/kwul0208/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {

	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listem: %v", err)
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)
	NewGRPCHandler(grpcServer)

	svc.CreateProduct(context.Background())

	log.Println("GRPC server started at ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}

	// r := gin.Default()

	// initializers.LoadEnvVariables()
	// db := initializers.ConnectDatabase()

	// productRepository := repository.NewProductRepository(db)
	// productUseCase := use_case.NewProductUseCase(productRepository)
	// productHandler := handler.NewProductHandler(productUseCase)

	// r.GET("/", productHandler.Index)
	// r.GET("/:id", productHandler.Show)
	// r.POST("/", productHandler.Create)
	// r.PUT("/:id", productHandler.Update)
	// r.DELETE("/", productHandler.Delete)

	// r.Run()
}
