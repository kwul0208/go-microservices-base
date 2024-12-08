package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/kwul0208/common/api"
)

type userHandler struct {
	client pb.UserServiceClient
}

func NewUserHandler(client pb.UserServiceClient) *userHandler {
	return &userHandler{client}
}

func (h *userHandler) HandleRegister(c *gin.Context) {
	var user pb.RegisterUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	// Call the gRPC service to create the product
	_, err := h.client.Create(c.Request.Context(), &pb.RegisterUserRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create user",
			"error":   err.Error(),
		})
		return
	}

	// Send success response
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Product created successfully",
	})
}

func (h *userHandler) HandlerLogin(c *gin.Context) {
	var user pb.LoginUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	// Call the gRPC service to create the product
	res, err := h.client.Login(c.Request.Context(), &pb.LoginUserRequest{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to login",
			"error":   err.Error(),
		})
		return
	}

	// Send success response
	c.JSON(http.StatusOK, gin.H{
		"result": res.Result,
	})
}
