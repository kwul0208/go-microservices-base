package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwul0208/user/request"
	"github.com/kwul0208/user/use_case"
)

type authHandler struct {
	authUseCase use_case.AuthUseCase
}

func NewAuthHandler(authUseCase use_case.AuthUseCase) *authHandler {
	return &authHandler{authUseCase}
}

func (h *authHandler) Create(c *gin.Context) {
	log.Printf("Hhhhhh2")

	var registerRequest request.RegisterRequest

	err := c.ShouldBindJSON(&registerRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := h.authUseCase.Create(registerRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "success",
		"data":    user,
	})
}

func (h *authHandler) Login(c *gin.Context) {
	var loginRequest request.LoginRequest

	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid input", "error": err.Error()})
		return
	}

	result, err := h.authUseCase.Login(loginRequest)
	if err != nil {
		if result["status"] == "error" {
			if result["message"] == "email or password are wrong" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, result) // 401 Unauthorized untuk kredensial yang salah
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, result) // 500 Internal Server Error untuk kesalahan lainnya
			}
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, result) // 400 Bad Request untuk kasus lainnya
		}
		return
	}

	c.JSON(http.StatusOK, result)
}
