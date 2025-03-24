package internal

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostCreateRequest struct {
	Title    string
	Content  string
	Category string
}

type PostSchema struct {
	ID        bson.ObjectID `bson:"_id"`
	Title     string        `bson:"title,omitempty"`
	Content   string        `bson:"content,omitempty"`
	LikeCount int           `bson:"like_count,omitempty"`
	Category  string        `bson:"category,omitempty"`
	CreatedAt time.Time     `bson:"create_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}
