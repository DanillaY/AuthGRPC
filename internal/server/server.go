package server

import (
	"context"
	Auth "main/internal/code_gen/proto_models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	Auth.UnimplementedAuthServer
	auth IAuth
}

type IAuth interface {
	Login(ctx context.Context, email string, gender string, phoneNumber string, password string, appID int64) (token string, err error)
	Register(ctx context.Context, email string, gender string, phoneNumber string, password string) (userID int64, err error)
}

func RegisterServer(grpc *grpc.Server, auth IAuth) {
	Auth.RegisterAuthServer(grpc, &Server{auth: auth})
}

func (s *Server) Login(ctx context.Context, loginR *Auth.LoginRequest) (*Auth.LoginResponse, error) {

	if loginR.GetEmail() == "" || loginR.GetGender() == "" || loginR.GetPassword() == "" || loginR.GetPhoneNumber() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid parameters")
	}
	token, err := s.auth.Login(ctx, loginR.GetEmail(), loginR.GetGender(), loginR.GetPhoneNumber(), loginR.GetPassword(), loginR.GetAppId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "User not found")
	}

	return &Auth.LoginResponse{Jwt: token}, nil
}
func (s *Server) Register(ctx context.Context, regR *Auth.RegRequest) (*Auth.RegResponse, error) {

	if regR.GetEmail() == "" || regR.GetGender() == "" || regR.GetPassword() == "" || regR.PhoneNumber == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid parameters")
	}
	userID, err := s.auth.Register(ctx, regR.GetEmail(), regR.GetGender(), regR.GetPhoneNumber(), regR.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while registering new user")
	}
	return &Auth.RegResponse{UserId: userID}, nil
}
