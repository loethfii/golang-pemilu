package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"luthfi/pemilu/domain"
)

type userRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) RegisterUser(ctx context.Context, user domain.User) (domain.User, error) {
	existingData := domain.User{}
	checkDB := r.DB.WithContext(ctx).Where("username = ?", user.Username).First(&existingData)
	if checkDB.RowsAffected > 0 {
		return domain.User{}, errors.New("username already exist")
	}
	
	err := r.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	
	return user, nil
}

func (r *userRepository) LoginUser(ctx context.Context, username, password string) (domain.User, error) {
	var user domain.User
	err := r.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	
	return user, nil
}
