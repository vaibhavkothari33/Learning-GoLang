package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavkothari33/backendapi/config"
	"github.com/vaibhavkothari33/backendapi/models"
	"github.com/vaibhavkothari33/backendapi/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection = config.GetCollection("users")

// register

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if user already exists
	var existingUser models.User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	} else if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	verificationCode := utils.GenerateVerificationCode()

	// Create new user
	user.ID = primitive.NewObjectID()
	user.Password = hashedPassword
	user.IsVerified = false
	user.VerificationCode = verificationCode
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register new user"})
		return
	}

	// Send verification email (non-blocking)
	go func() {
		if err := utils.SendVerificationEmail(user.Email, verificationCode); err != nil {
			fmt.Println("Failed to send verification email:", err.Error())
		}
	}()

	// Sanitize user before returning
	user.Password = ""
	user.VerificationCode = ""

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully. Please check your email for the verification code.",
		"user":    user,
	})
}

func VerifyEmail(c *gin.Context) {
	var req models.VerifyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"email":             req.Email,
		"verification_code": req.Code,
		"is_verified":       false,
	}

	update := bson.M{
		"$set": bson.M{
			"is_verified": true,
			"updated_at":  time.Now(),
		},
		"$unset": bson.M{
			"verification_code": "",
		},
	}
	result := userCollection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code ot email"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify email"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email Verified successfully"})
}

func Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentails"})
		return
	}

	if !user.IsVerified {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please verify you email"})
		return
	}

	if err := utils.CheckPassword(user.Password, req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credential"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	user.Password = ""
	user.VerificationCode = ""

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful",
		"token":   token,
		"user":    user,
	})
}

func ForgotPassword(c *gin.Context) {
	var req models.ForgotPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if user exists
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		// Don't reveal if email exists or not for security
		c.JSON(http.StatusOK, gin.H{"message": "If email exists, password reset link has been sent"})
		return
	}

	// Generate reset token
	resetToken := primitive.NewObjectID().Hex()
	resetTokenExpiry := time.Now().Add(1 * time.Hour)

	// Update user with reset token
	update := bson.M{
		"$set": bson.M{
			"reset_token":        resetToken,
			"reset_token_expiry": resetTokenExpiry,
			"updated_at":         time.Now(),
		},
	}

	_, err = userCollection.UpdateOne(ctx, bson.M{"email": req.Email}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
		return
	}

	// Send reset email
	go func() {
		if err := utils.SendPasswordResetEmail(user.Email, resetToken); err != nil {
			println("Failed to send password reset email:", err.Error())
		}
	}()

	c.JSON(http.StatusOK, gin.H{"message": "If email exists, password reset link has been sent"})
}

// ResetPassword resets user password with token
func ResetPassword(c *gin.Context) {
	var req models.ResetPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find user with valid reset token
	filter := bson.M{
		"reset_token": req.Token,
		"reset_token_expiry": bson.M{
			"$gt": time.Now(),
		},
	}

	var user models.User
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired reset token"})
		return
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Update password and remove reset token
	update := bson.M{
		"$set": bson.M{
			"password":   hashedPassword,
			"updated_at": time.Now(),
		},
		"$unset": bson.M{
			"reset_token":        "",
			"reset_token_expiry": "",
		},
	}

	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": user.ID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}
