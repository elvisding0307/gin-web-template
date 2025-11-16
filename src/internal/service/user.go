package service

import (
	"gin-web-template/internal/dao"
	"gin-web-template/internal/errors"
	"gin-web-template/internal/model"
)

func RegisterService(username, password string) (uint64, errors.SrvErr) {
	db, err := dao.GetMysqlInstance()
	if err != nil {
		return 0, errors.ErrDatabaseConnection
	}

	// Check if user already exists
	var existingUser model.User
	if err := db.Where("username = ?", username).First(&existingUser).Error; err == nil {
		// User already exists
		return 0, errors.ErrUserAlreadyExists
	}

	// Create new user
	newUser := model.User{
		Username: username,
		Password: password,
	}

	if err := db.Create(&newUser).Error; err != nil {
		return 0, errors.ErrRegistrationFailed
	}

	return newUser.UserId, nil
}

func ChangePasswordService(userId uint64, oldPassword, newPassword string) errors.SrvErr {
	db, err := dao.GetMysqlInstance()
	if err != nil {
		return errors.ErrDatabaseConnection
	}

	// Find user by user ID
	var user model.User
	if err := db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return errors.ErrUserNotFound
	}

	// Check if old password matches
	if user.Password != oldPassword {
		return errors.ErrInvalidCredentials
	}

	// Update password
	user.Password = newPassword
	if err := db.Save(&user).Error; err != nil {
		return errors.ErrRegistrationFailed
	}

	return nil
}
