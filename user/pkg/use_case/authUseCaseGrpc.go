package use_case

import (
	"errors"

	models "github.com/kwul0208/user/pkg/model"
	"github.com/kwul0208/user/pkg/pb"
	"github.com/kwul0208/user/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCaseGrpc interface {
	Register(request *pb.RegisterRequest) (models.User, error)
	Login(request *pb.LoginRequest) (models.User, error)
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

func (au *authUseCaseGrpc) Login(grpcReq *pb.LoginRequest) (models.User, error) {
	user, err := au.authRepository.FindByEmail(grpcReq.Email)
	if err != nil {
		return models.User{}, errors.New("email or password are wrong")
	}

	// Jika user tidak ditemukan
	if user.Id == 0 {
		return models.User{}, errors.New("email or password are wrong")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(grpcReq.Password))
	if err != nil {
		return models.User{}, errors.New("email or password are wrong")
	}

	// Return response
	return user, nil
}
