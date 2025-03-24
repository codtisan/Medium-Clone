package internal

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserSignupRequest struct {
	Username string
	Password string
	Email    string
}

type UserLoginRequest struct {
	Username string
	Password string
}

type UserProfile struct {
	ID        bson.ObjectID `bson:"_id"`
	Username  string        `bson:"username,omitempty"`
	Password  string        `bson:"password,omitempty"`
	Email     string        `bson:"email,omitempty"`
	CreatedAt time.Time     `bson:"create_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}
