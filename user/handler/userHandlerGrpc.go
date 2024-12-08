package handler

import (
	"context"
	"log"

	pb "github.com/kwul0208/common/api"
	"github.com/kwul0208/user/use_case"
	"google.golang.org/grpc"
)

type authHandlerGrpc struct {
	pb.UnimplementedUserServiceServer
	user_guc use_case.AuthUseCaseGrpc
}

func NewAuthGrpcHandler(grpcServer *grpc.Server, userUseCase use_case.AuthUseCaseGrpc) {
	handler := &authHandlerGrpc{
		user_guc: userUseCase,
	}

	pb.RegisterUserServiceServer(grpcServer, handler)
}

func (h *authHandlerGrpc) Create(ctx context.Context, rr *pb.RegisterUserRequest) (*pb.User, error) {

	data := &pb.RegisterUserRequest{
		Name:     rr.Name,
		Email:    rr.Email,
		Password: rr.Password,
	}

	_, err := h.user_guc.Create(data)

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	response := &pb.User{
		Name:     data.Name, // Populate fields as per your `User` definition
		Email:    data.Email,
		Password: data.Password,
	}

	return response, nil
}

func (h *authHandlerGrpc) Login(ctx context.Context, lr *pb.LoginUserRequest) (*pb.LoginUserResponseWrapper, error) {
	data := &pb.LoginUserRequest{
		Email:    lr.Email,
		Password: lr.Password,
	}
	user, err := h.user_guc.Login(data)
	if err != nil {
		return &pb.LoginUserResponseWrapper{
			Result: &pb.LoginUserResponseWrapper_Failed{
				Failed: &pb.LoginUserResponseFailed{
					Message: err.Error(),
				},
			},
		}, nil
	}

	return &pb.LoginUserResponseWrapper{
		Result: &pb.LoginUserResponseWrapper_Success{
			Success: &pb.LoginUserResponse{
				Token: user.Token, // Replace with actual token
				User: &pb.User{
					Name:  user.User.Name,
					Email: user.User.Email,
				},
			},
		},
	}, nil
}
