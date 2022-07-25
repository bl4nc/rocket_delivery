package mongo

import (
	"time"

	"github.com/bl4nc/rocket_delivery/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser()
}

type userRepository struct {
	client     *mongo.Client
	database   string
	collection string
	timeout    time.Duration
}

func newUserRepository(client *mongo.Client, database, collection string, timeout time.Duration) UserRepository {
	return &userRepository{
		client:     client,
		database:   database,
		collection: collection,
		timeout:    timeout,
	}
}

func (u *userRepository) CreateUser(newUser *entities.User) {
	ctx, cancel := u.getContext()
	defer cancel()

}
