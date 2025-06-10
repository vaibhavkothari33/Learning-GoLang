package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username         string             `json:"username" bson:"username" binding:"required"`
	Email            string             `json:"email" bson:"email" binding:"required,email"`
	Password         string             `json:"password,omitempty" bson:"password" binding:"required"`
	IsVerified       bool               `json:"is_verified" bson:"is_verified"`
	VerificationCode string             `json:"-" bson:"verification_code,omitempty"`
	ResetToken       string             `json:"-" bson:"reset_token,omitempty"`
	ResetTokenExpiry time.Time          `json:"-" bson:"reset_token_expiry,omitempty"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" bson:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type VerifyRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}
