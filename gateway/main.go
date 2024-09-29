package main

import (
	"log"
	"net/http"

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

	mux := http.NewServeMux()
	handlerInit := handler.NewHandler(c)
	handlerInit.RegisterRoutes(mux)

	log.Printf("Starting server gateway at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start gateway")
	}
}
