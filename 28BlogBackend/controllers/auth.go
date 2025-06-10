package controllers

import (
	"context"
	"fmt"
	"githconfig/vaibhavkothari33/backendapi/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	result := userCollection.FindOneAndUpdate(ctx,filter,update)
	if result.Err() !=nil{
		if result.Err() == mongo.ErrNoDocuments{
			c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid verification code ot email"})
		} else{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to verify email"})
		}
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"Email Verified successfully"})
}



// package controllers

// import (
//     "blog-backend/config"
//     "blog-backend/models"
//     "blog-backend/utils"
//     "context"
//     "net/http"
//     "time"

//     "github.com/gin-gonic/gin"
//     "go.mongodb.org/mongo-driver/bson"
//     "go.mongodb.org/mongo-driver/bson/primitive"
//     "go.mongodb.org/mongo-driver/mongo"
// )

// var userCollection = config.GetCollection("users")

// // Register creates a new user account
// func Register(c *gin.Context) {
//     var user models.User

//     // Bind JSON to user struct
//     if err := c.ShouldBindJSON(&user); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     // Check if user already exists
//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

//     var existingUser models.User
//     err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
//     if err == nil {
//         c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
//         return
//     }

//     // Hash password
//     hashedPassword, err := utils.HashPassword(user.Password)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
//         return
//     }

//     // Generate verification code
//     verificationCode := utils.GenerateVerificationCode()

//     // Create user object
//     user.ID = primitive.NewObjectID()
//     user.Password = hashedPassword
//     user.IsVerified = false
//     user.VerificationCode = verificationCode
//     user.CreatedAt = time.Now()
//     user.UpdatedAt = time.Now()

//     // Insert user into database
//     _, err = userCollection.InsertOne(ctx, user)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
//         return
//     }

//     // Send verification email
//     go func() {
//         if err := utils.SendVerificationEmail(user.Email, verificationCode); err != nil {
//             // Log error but don't fail the request
//             // In production, you might want to use a proper logging system
//             println("Failed to send verification email:", err.Error())
//         }
//     }()

//     // Remove sensitive data before responding
//     user.Password = ""
//     user.VerificationCode = ""

//     c.JSON(http.StatusCreated, gin.H{
//         "message": "User created successfully. Please check your email for verification code.",
//         "user":    user,
//     })
// }

// // VerifyEmail verifies user's email with the code
// func VerifyEmail(c *gin.Context) {
//     var req models.VerifyRequest

//     if err := c.ShouldBindJSON(&req); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

//     // Find user and verify code
//     filter := bson.M{
//         "email":             req.Email,
//         "verification_code": req.Code,
//         "is_verified":       false,
//     }

//     update := bson.M{
//         "$set": bson.M{
//             "is_verified":       true,
//             "updated_at":        time.Now(),
//         },
//         "$unset": bson.M{
//             "verification_code": "",
//         },
//     }

//     result := userCollection.FindOneAndUpdate(ctx, filter, update)
//     if result.Err() != nil {
//         if result.Err() == mongo.ErrNoDocuments {
//             c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code or email"})
//         } else {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify email"})
//         }
//         return
//     }

//     c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
// }

// // Login authenticates user and returns JWT token
// func Login(c *gin.Context) {
//     var req models.LoginRequest

//     if err := c.ShouldBindJSON(&req); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

//     // Find user by email
//     var user models.User
//     err := userCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
//     if err != nil {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
//         return
//     }

//     // Check if email is verified
//     if !user.IsVerified {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "Please verify your email first"})
//         return
//     }

//     // Check password
//     if err := utils.CheckPassword(user.Password, req.Password); err != nil {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
//         return
//     }

//     // Generate JWT token
//     token, err := utils.GenerateToken(user.ID, user.Email, user.Username)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
//         return
//     }

//     // Remove sensitive data
//     user.Password = ""
//     user.VerificationCode = ""

//     c.JSON(http.StatusOK, gin.H{
//         "message": "Login successful",
//         "token":   token,
//         "user":    user,
//     })
// }

// // ForgotPassword sends password reset email
// func ForgotPassword(c *gin.Context) {
//     var req models.ForgotPasswordRequest

//     if err := c.ShouldBindJSON(&req); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

//     // Check if user exists
//     var user models.User
//     err := userCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
//     if err != nil {
//         // Don't reveal if email exists or not for security
//         c.JSON(http.StatusOK, gin.H{"message": "If email exists, password reset link has been sent"})
//         return
//     }

//     // Generate reset token
//     resetToken := primitive.NewObjectID().Hex()
//     resetTokenExpiry := time.Now().Add(1 * time.Hour)

//     // Update user with reset token
//     update := bson.M{
//         "$set": bson.M{
//             "reset_token":        resetToken,
//             "reset_token_expiry": resetTokenExpiry,
//             "updated_at":         time.Now(),
//         },
//     }

//     _, err = userCollection.UpdateOne(ctx, bson.M{"email": req.Email}, update)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
//         return
//     }

//     // Send reset email
//     go func() {
//         if err := utils.SendPasswordResetEmail(user.Email, resetToken); err != nil {
//             println("Failed to send password reset email:", err.Error())
//         }
//     }()

//     c.JSON(http.StatusOK, gin.H{"message": "If email exists, password reset link has been sent"})
// }

// // ResetPassword resets user password with token
// func ResetPassword(c *gin.Context) {
//     var req models.ResetPasswordRequest

//     if err := c.ShouldBindJSON(&req); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

//     // Find user with valid reset token
//     filter := bson.M{
//         "reset_token": req.Token,
//         "reset_token_expiry": bson.M{
//             "$gt": time.Now(),
//         },
//     }

//     var user models.User
//     err := userCollection.FindOne(ctx, filter).Decode(&user)
//     if err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired reset token"})
//         return
//     }

//     // Hash new password
//     hashedPassword, err := utils.HashPassword(req.NewPassword)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
//         return
//     }

//     // Update password and remove reset token
//     update := bson.M{
//         "$set": bson.M{
//             "password":   hashedPassword,
//             "updated_at": time.Now(),
//         },
//         "$unset": bson.M{
//             "reset_token":        "",
//             "reset_token_expiry": "",
//         },
//     }

//     _, err = userCollection.UpdateOne(ctx, bson.M{"_id": user.ID}, update)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
//         return
//     }

//     c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
// }

// // Logout (client-side token removal, but we can add token blacklisting here)
// func Logout(c *gin.Context) {
//     // In a simple JWT implementation, logout is handled client-side by removing the token
//     // For more security, you could implement token blacklisting here
//     c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
// }
// ```
