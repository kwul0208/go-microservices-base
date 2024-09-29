package main

import (
	"github.com/gin-gonic/gin"
)

func init() {
}

func main() {
	r := gin.Default()

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

	r.Run()
}
