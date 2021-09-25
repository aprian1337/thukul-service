package users

import (
	"aprian1337/thukul-service/app/middlewares"
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/utilities"
	"aprian1337/thukul-service/utilities/constants"
	"context"
	"errors"
	"log"
	"time"
)

type UserUsecase struct {
	Repo    Repository
	Timeout time.Duration
	jwtAuth *middlewares.ConfigJWT
}

func NewUserUsecase(repo Repository, timeout time.Duration, jwtAuth *middlewares.ConfigJWT) *UserUsecase {
	return &UserUsecase{
		Repo:    repo,
		Timeout: timeout,
		jwtAuth: jwtAuth,
	}
}

func (uc *UserUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) GetById(ctx context.Context, id uint) (Domain, error) {
	user, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
func (uc *UserUsecase) Create(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("email is required")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("password is required")
	}

	_, err := time.Parse(constants.BirthdayFormat, domain.Birthday)
	if err != nil {
		return Domain{}, businesses.ErrInvalidDate
	}

	domain.Password, _ = utilities.HashPassword(domain.Password)

	user, err := uc.Repo.Create(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (Domain, string, error) {
	if email == "" || password == "" {
		return Domain{}, "", businesses.ErrUsernamePasswordNotFound
	}

	user, err := uc.Repo.GetByEmail(ctx, email)

	if err != nil {
		return Domain{}, "", err
	}

	if !utilities.CheckPassword(password, user.Password) {
		return Domain{}, "", businesses.ErrInvalidAuthentication
	}

	//JWT
	token, errToken := uc.jwtAuth.GenerateTokenJWT(user.ID)
	if errToken != nil {
		log.Println(errToken)
	}
	if token == "" {
		return Domain{}, "", businesses.ErrInvalidAuthentication
	}

	return user, token, nil
}
