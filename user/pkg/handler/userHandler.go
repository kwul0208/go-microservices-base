package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/kwul0208/user/pkg/db"
	models "github.com/kwul0208/user/pkg/model"
	"github.com/kwul0208/user/pkg/pb"
	"github.com/kwul0208/user/pkg/use_case"
	"github.com/kwul0208/user/pkg/utils"
)

type Server struct {
	H        db.Handler
	Jwt      utils.JWTWrapper
	User_guc use_case.AuthUseCaseGrpc
	pb.UnimplementedUserServiceServer
}

func (h *Server) Register(ctx context.Context, rr *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	if h.User_guc == nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  "AuthUseCaseGrpc is not initialized",
		}, errors.New("nil AuthUseCaseGrpc instance")
	}

	data := &pb.RegisterRequest{
		Email:    rr.Email,
		Password: rr.Password,
	}

	_, err := h.User_guc.Register(data)

	if err != nil {
		// log.Printf("Error creating user: %v", err)
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (h *Server) Login(ctx context.Context, lr *pb.LoginRequest) (*pb.LoginResponse, error) {
	data := &pb.LoginRequest{
		Email:    lr.Email,
		Password: lr.Password,
	}
	user, err := h.User_guc.Login(data)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	token, _ := h.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (h *Server) Validate(ctx context.Context, vr *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	fmt.Println("Auth Service: Validate")
	claims, err := h.Jwt.ValidateToken(vr.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User

	if result := h.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}
