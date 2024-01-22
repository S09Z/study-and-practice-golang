package datastore

import (
	"context"
	"errors"
	"log"
	"sync"

	pb "github.com/yourusername/user-service/pb"
)

type Datastore struct {
	mu    sync.RWMutex
	users map[string]*pb.User
}

func NewDatastore() *Datastore {
	return &Datastore{
		users: make(map[string]*pb.User),
	}
}

func (d *Datastore) CreateUser(ctx context.Context, user *pb.User) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Check if the user already exists
	if _, ok := d.users[user.GetId()]; ok {
		return errors.New("user already exists")
	}

	// Add the user to the map
	d.users[user.GetId()] = user
	log.Printf("User created: %v", user)
	return nil
}

// Implement other CRUD methods (GetUser, UpdateUser, DeleteUser)
