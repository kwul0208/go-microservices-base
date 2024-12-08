package use_case

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	models "github.com/kwul0208/user/model"
	"github.com/kwul0208/user/repository"
	"github.com/kwul0208/user/request"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	Create(request request.RegisterRequest) (models.User, error)
	Login(request request.LoginRequest) (gin.H, error)
}

type authUseCase struct {
	authRepository repository.AuthRepository
}

func NewAuthUseCase(authRepository repository.AuthRepository) *authUseCase {
	return &authUseCase{authRepository}
}

func (au *authUseCase) Create(registerRequest request.RegisterRequest) (models.User, error) {
	log.Printf("Hhhhhh")
	hash, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 10)
	if err != nil {
		// log.Panic(err.Error())
	}

	user := models.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: string(hash),
		Role:     registerRequest.Role,
	}

	newUser, err := au.authRepository.Create(user)

	return newUser, err
}

func (au *authUseCase) Login(loginRequest request.LoginRequest) (gin.H, error) {
	user, err := au.authRepository.FindByEmail(loginRequest.Email)
	if err != nil {
		return gin.H{
			"status":  "error",
			"message": "Failed to retrieve user data",
		}, err
	}

	// Jika user tidak ditemukan
	if user.Id == 0 {
		return gin.H{
			"status":  "error",
			"message": "email or password are wrong",
		}, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return gin.H{
			"status":  "error",
			"message": "email or password are wrong",
		}, errors.New("password mismatch")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Sign dan dapatkan token string
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return gin.H{
			"status":  "error",
			"message": "Failed to generate token",
		}, err
	}

	// Kembalikan token
	return gin.H{
		"status":  "success",
		"message": "Login successful",
		"token":   tokenString,
	}, nil
}
