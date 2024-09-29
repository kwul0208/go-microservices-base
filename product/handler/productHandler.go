package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kwul0208/product/request"
	"github.com/kwul0208/product/use_case"
)

type productHandler struct {
	productUseCase use_case.ProductUseCase
}

func NewProductHandler(productUseCase use_case.ProductUseCase) *productHandler {
	return &productHandler{productUseCase}
}

func (h *productHandler) Index(c *gin.Context) {
	products, err := h.productUseCase.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}
func (h *productHandler) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.productUseCase.GetById(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})

}

func (h *productHandler) Create(c *gin.Context) {
	var productRequest request.ProductRequest

	err := c.ShouldBindJSON(&productRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	product, err := h.productUseCase.Create(productRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "success",
		"data":    product,
	})
}

func (h *productHandler) Update(c *gin.Context) {
	var productRequest request.ProductRequest
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&productRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	}

	updatedProduct, err := h.productUseCase.Update(id, productRequest)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": err.Error(),
		"data":    updatedProduct,
	})

}
func (h *productHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.productUseCase.Delete(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	}

	c.AbortWithStatusJSON(http.StatusAccepted, gin.H{"message": "success delete data", "pr": product})
}
