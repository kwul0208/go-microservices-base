package use_case

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	pb "github.com/kwul0208/common/api"
	models "github.com/kwul0208/user/model"
	"github.com/kwul0208/user/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCaseGrpc interface {
	Create(request *pb.RegisterUserRequest) (models.User, error)
	Login(request *pb.LoginUserRequest) (*pb.LoginUserResponse, error)
}

type authUseCaseGrpc struct {
	authRepository repository.AuthRepository
}

func NewAuthUseCaseGrpc(authRepository repository.AuthRepository) *authUseCaseGrpc {
	return &authUseCaseGrpc{authRepository}
}

func (au *authUseCaseGrpc) Create(grpcReq *pb.RegisterUserRequest) (models.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(grpcReq.Password), 10)
	if err != nil {
	}

	user := models.User{
		Name:     grpcReq.Name,
		Email:    grpcReq.Email,
		Password: string(hash),
	}

	newUser, err := au.authRepository.Create(user)

	return newUser, err
}

func (au *authUseCaseGrpc) Login(grpcReq *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
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
	return &pb.LoginUserResponse{
		Token: tokenString,
		User: &pb.User{
			ID:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
