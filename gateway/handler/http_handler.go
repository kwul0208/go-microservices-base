package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "github.com/kwul0208/common/api"
)

type handler struct {
	client pb.ProductServiceClient
}

func NewHandler(client pb.ProductServiceClient) *handler {
	return &handler{client}
}

func (h *handler) HandlerGetProduct(c *gin.Context) {
	product, _ := h.client.GetProducts(c.Request.Context(), &pb.Empty{})

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "getting product successfully",
		"data":    product,
	})
}

func (h *handler) HandlerGetProductById(c *gin.Context) {
	// Get the "id" parameter from the request URL
	idStr := c.Query("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Failed to parse id",
		})
		return
	}

	// Call gRPC to get the product by ID
	product, err := h.client.GetProductById(c.Request.Context(), &pb.ProductID{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to get product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "getting product detail successfully",
		"data":    product,
	})
}

func (h *handler) HandleCreateProduct(c *gin.Context) {
	// Bind the incoming JSON request to a `ProductOnly` struct
	var product pb.ProductOnly
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	// Call the gRPC service to create the product
	_, err := h.client.CreateProduct(c.Request.Context(), &pb.CreateProductRequest{
		ProductOnly: &product,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create product",
		})
		return
	}

	// Send success response
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Product created successfully",
	})
}

func (h *handler) HandleUpdateProduct(c *gin.Context) {
	// Parse the "id" query parameter from the request
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Failed to parse id",
		})
		return
	}

	// Bind the JSON request body to the ProductOnly struct
	var product pb.ProductOnly
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	// Call the gRPC service to update the product
	_, err = h.client.UpdateProduct(c.Request.Context(), &pb.UpdateProductRequest{
		ID:          id,
		ProductOnly: &product,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to update product",
		})
		return
	}

	// Send success response
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Product updated successfully",
	})
}

func (h *handler) HandleDeleteProduct(c *gin.Context) {
	// Get the ID from the query parameter
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Failed to parse id",
		})
		return
	}

	// Call the gRPC service to delete the product
	product, err := h.client.DeleteProduct(c.Request.Context(), &pb.ProductID{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to delete product",
		})
		return
	}

	// Send success response
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Product deleted successfully",
		"data":    product,
	})
}
