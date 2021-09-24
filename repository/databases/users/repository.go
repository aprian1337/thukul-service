package users

import (
	"aprian1337/thukul-service/business/users"
	"aprian1337/thukul-service/deliveries/users/requests"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type PostgresUserRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresUserRepository(conn *gorm.DB) users.Repository {
	return &PostgresUserRepository{
		ConnPostgres: conn,
	}
}

func (repo *PostgresUserRepository) GetAll(ctx context.Context) ([]users.Domain, error) {
	var user []Users
	err := repo.ConnPostgres.Find(&Users{})
	if err.Error != nil {
		return []users.Domain{}, err.Error
	}
	return ToListDomain(user), nil
}

func (repo *PostgresUserRepository) GetById(id uint, ctx context.Context) (users.Domain, error) {
	var user Users
	err := repo.ConnPostgres.Find(&user, "id = ?", id)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.ToDomain(), nil
}

func (repo *PostgresUserRepository) Create(ctx context.Context, register requests.UserRegister) (users.Domain, error) {
	birthday, errTime := time.Parse("2006-01-02", register.Birthday)
	if errTime != nil {
		return users.Domain{}, errors.New("invalid birthday (must yyyy-mm-dd)")
	}
	user := Users{
		SalaryId: register.SalaryId,
		Name:     register.Name,
		Password: register.Password,
		Email:    register.Email,
		Phone:    register.Phone,
		Gender:   register.Gender,
		Birthday: birthday,
		Address:  register.Address,
		Company:  register.Company,
	}
	err := repo.ConnPostgres.Create(&user)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.ToDomain(), nil
}

func (repo *PostgresUserRepository) Login(ctx context.Context, login requests.UserLogin) (users.Domain, error) {
	var user Users
	err := repo.ConnPostgres.Find(&user, "email = ? AND password = ?", login.Email, login.Password)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.ToDomain(), nil
}
