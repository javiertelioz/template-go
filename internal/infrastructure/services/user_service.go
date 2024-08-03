package services

import (
	"context"
	pb "github.com/javiertelioz/template-clean-architecture-go/internal/infrastructure/proto/user"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	// Aquí puedes agregar dependencias como repositorios, etc.
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// Lógica para crear un usuario
	return &pb.CreateUserResponse{Id: "generated-id"}, nil
}

func (s *UserServiceServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	// Lógica para obtener un usuario por ID
	return &pb.GetUserByIdResponse{Id: req.Id, Name: "example", Email: "example@example.com"}, nil
}
