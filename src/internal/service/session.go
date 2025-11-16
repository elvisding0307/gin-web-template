// internal/service/session.go
package service

import (
	"gin-web-template/internal/config"
	"gin-web-template/internal/dao"
	"gin-web-template/internal/errors"
	"gin-web-template/internal/model"
	"time"

	"github.com/golang-jwt/jwt"
)

// LoginService authenticates a user against the database
func LoginService(username, password string) (interface{}, errors.SrvErr) {
	// Get database instance
	db, err := dao.GetMysqlInstance()
	if err != nil {
		return nil, errors.ErrDatabaseConnection
	}

	// Find user by username
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		// We return a generic error to prevent username enumeration attacks
		return nil, errors.ErrUserNotFound
	}

	// Compare hashed password
	if user.Password != password {
		return nil, errors.ErrInvalidCredentials
	}

	// Generate a simple token (in production, use JWT or proper session management)
	token, err := generateToken(user.UserId)

	if err != nil {
		return nil, errors.ErrTokenGeneration
	}

	// Return user data without sensitive information
	return map[string]interface{}{
		"user_id":  user.UserId,
		"username": user.Username,
		"token":    token,
	}, nil
}

// Helper function to generate a simple token
func generateToken(userID uint64) (string, error) {
	// Get server configuration
	cfg, err := config.ServerConfig()
	if err != nil {
		return "", err
	}

	// Create JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":     time.Now().Unix(),                     // Issued at time
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(cfg.ServerSecretKey))
	if err != nil {
		return "", errors.ErrTokenGeneration
	}

	return tokenString, nil
}
