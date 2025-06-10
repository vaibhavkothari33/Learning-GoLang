package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title" binding:"required"`
	Content   string             `json:"content" bson:"content" binding:"required"`
	Author    string             `json:"author" bson:"author"`
	AuthorID  primitive.ObjectID `json:"author_id" bson:"author_id"`
	Tags      []string           `json:"tags" bson:"tags"`
	Published bool               `json:"published" bson:"published"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CreateBlogRequest struct {
	Title     string   `json:"title" binding:"required"`
	Content   string   `json:"content" binding:"required"`
	Tags      []string `json:"tags"`
	Published bool     `json:"published"`
}

type UpdateBlogRequest struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	Published *bool    `json:"published"`
}
