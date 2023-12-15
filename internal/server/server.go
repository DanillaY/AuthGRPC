package server

import (
	"context"
	Auth "main/internal/code_gen/proto_models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IAuth interface {
	Login(ctx context.Context, email string, gender string, phoneNumber string, password string) (token string, err error)
	Register(ctx context.Context, email string, gender string, phoneNumber string, password string) (userID int64, err error)
}

func RegisterServer(grpc *grpc.Server, auth IAuth) {
	Auth.RegisterAuthServer(grpc, &Server{auth: auth})
}

func (s *Server) Login(ctx context.Context, loginR *Auth.LoginRequest) (*Auth.LoginResponse, error) {
	if loginR.GetEmail() == "" || loginR.GetGender() == "" || loginR.GetPassword() == "" || loginR.PhoneNumber == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid parameters")
	}
	//TODO: заменить на реальный jwt токен
	return &Auth.LoginResponse{Jwt: "test"}, nil
}
func (s *Server) Register(ctx context.Context, regR *Auth.RegRequest) (*Auth.RegResponse, error) {
	if regR.GetEmail() == "" || regR.GetGender() == "" || regR.GetPassword() == "" || regR.PhoneNumber == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid parameters")
	}
	//TODO: заменить на id пользователя
	return &Auth.RegResponse{UserId: 123}, nil
}

type Server struct {
	Auth.UnimplementedAuthServer
	auth IAuth
}
