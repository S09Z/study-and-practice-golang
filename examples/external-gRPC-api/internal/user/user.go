package user

import (
	"context"

	"github.com/yourusername/user-service/internal/datastore"
	pb "github.com/yourusername/user-service/pb"
)

type UserManager struct {
	ds *datastore.Datastore
}

func NewUserManager(ds *datastore.Datastore) *UserManager {
	return &UserManager{
		ds: ds,
	}
}

func (um *UserManager) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := req.GetUser()
	err := um.ds.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{User: user}, nil
}

// Implement other CRUD methods (GetUser, UpdateUser, DeleteUser)
