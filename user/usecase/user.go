package usecase

import (
	"context"
	"errors"
	"luthfi/pemilu/domain"
	"luthfi/pemilu/utils/hash"
)

type userUsecase struct {
	domain.UserRepository
}

func NewUserUseCase(ur domain.UserRepository) domain.UserUseCase {
	return &userUsecase{ur}
}

func (u *userUsecase) RegisterUser(user domain.User) (domain.User, error) {
	enPass, err := hash.PasswordHash(user.Password)
	if err != nil {
		return domain.User{}, err
	}
	user.Password = enPass
	
	ctx := context.Background()
	
	dataUser, err := u.UserRepository.RegisterUser(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	
	return dataUser, nil
}

func (u *userUsecase) LoginUser(username, password string) (domain.User, error) {
	ctx := context.Background()
	user, err := u.UserRepository.LoginUser(ctx, username, password)
	if err != nil {
		return domain.User{}, errors.New("invalid username or password")
	}
	
	isSuccess := hash.CheckPasswordHash(password, user.Password)
	if !isSuccess {
		return domain.User{}, errors.New("invalid username or password")
	}
	
	return user, nil
}
