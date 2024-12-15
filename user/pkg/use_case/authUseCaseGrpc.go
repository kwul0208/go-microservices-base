package use_case

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	models "github.com/kwul0208/user/pkg/model"
	"github.com/kwul0208/user/pkg/pb"
	"github.com/kwul0208/user/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCaseGrpc interface {
	Register(request *pb.RegisterRequest) (models.User, error)
	Login(request *pb.LoginRequest) (*pb.LoginResponse, error)
}

type authUseCaseGrpc struct {
	authRepository repository.AuthRepository
}

func NewAuthUseCaseGrpc(authRepository repository.AuthRepository) *authUseCaseGrpc {
	return &authUseCaseGrpc{authRepository}
}

func (au *authUseCaseGrpc) Register(grpcReq *pb.RegisterRequest) (models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(grpcReq.Password), 10)
	if err != nil {
	}

	user := models.User{
		Email:    grpcReq.Email,
		Password: string(hash),
	}

	newUser, err := au.authRepository.Register(user)

	return newUser, err
}

func (au *authUseCaseGrpc) Login(grpcReq *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := au.authRepository.FindByEmail(grpcReq.Email)
	if err != nil {
		return nil, errors.New("email or password are wrong")
	}

	// Jika user tidak ditemukan
	if user.Id == 0 {
		return nil, errors.New("email or password are wrong")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(grpcReq.Password))
	if err != nil {
		return nil, errors.New("email or password are wrong")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Sign dan dapatkan token string
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return nil, errors.New("Failed generate token")

	}

	// Return response
	return &pb.LoginResponse{
		Token: tokenString,
	}, err
}
