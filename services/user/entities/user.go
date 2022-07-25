package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:_id`
	CreatedAt time.Time          `bson:created_at`
	Name      string             `bson:name`
	City      string             `bson:city`
	State     string             `bson:state`
	Cep       string             `bson:cep`
	LastLogin string             `bson:last_login`
}
