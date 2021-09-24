package users

import (
	"aprian1337/thukul-service/deliveries/users/requests"
	"aprian1337/thukul-service/utilities"
	"context"
	"errors"
	"time"
)

type UserUsecase struct {
	Repo    Repository
	Timeout time.Duration
}

func NewUserUsecase(repo Repository, timeout time.Duration) *UserUsecase {
	return &UserUsecase{
		Repo:    repo,
		Timeout: timeout,
	}
}

func (uc *UserUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) GetById(id uint, ctx context.Context) (Domain, error) {
	user, err := uc.Repo.GetById(id, ctx)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
func (uc *UserUsecase) Create(ctx context.Context, register requests.UserRegister) (Domain, error) {
	if register.Email == "" {
		return Domain{}, errors.New("email is required")
	}

	if register.Password == "" {
		return Domain{}, errors.New("password is required")
	}
	register.Password, _ = utilities.HashPassword(register.Password)

	user, err := uc.Repo.Create(ctx, register)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
func (uc *UserUsecase) Login(ctx context.Context, login requests.UserLogin) (Domain, error) {
	if login.Email == "" {
		return Domain{}, errors.New("email is required")
	}

	if login.Password == "" {
		return Domain{}, errors.New("password is required")
	}

	user, err := uc.Repo.Login(ctx, login)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
