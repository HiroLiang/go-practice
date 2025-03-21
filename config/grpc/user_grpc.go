package grpc

import (
	"TestProject/db"
	pb "TestProject/proto"
	"context"
	"fmt"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserService) GetAllUser(ctx context.Context, req *pb.Empty) (*pb.UserListResponse, error) {
	users, err := db.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.UserResponse
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.UserResponse{
			Id:   int32(u.ID),
			Name: u.Name,
		})
	}

	return &pb.UserListResponse{Users: pbUsers}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := db.GetUserByID(req.Id)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return &pb.UserResponse{
		Id:   int32(user.ID),
		Name: user.Name,
	}, nil
}
