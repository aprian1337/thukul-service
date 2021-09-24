package users

import (
	"aprian1337/thukul-service/business/users"
	"aprian1337/thukul-service/deliveries/users/requests"
	"aprian1337/thukul-service/utilities"
	"context"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresUserRepository(conn *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		ConnPostgres: conn,
	}
}

func (repo *PostgresUserRepository) GetById(ctx context.Context, id uint) (users.Domain, error) {
	var user Users
	err := repo.ConnPostgres.Find(&user, "id = ?", id)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.ToDomain(), nil
}

func (repo *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (users.Domain, error) {
	var user Users
	err := repo.ConnPostgres.Find(&user, "email = ?", email)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.ToDomain(), nil
}

func (repo *PostgresUserRepository) Create(ctx context.Context, register *users.Domain) (users.Domain, error) {
	user := Users{
		SalaryId: register.SalaryId,
		Name:     register.Name,
		Password: register.Password,
		Email:    register.Email,
		Phone:    register.Phone,
		Gender:   register.Gender,
		Birthday: register.Birthday,
		Address:  register.Address,
		Company:  register.Company,
	}
	err := repo.ConnPostgres.Create(&user)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.ToDomain(), nil
}

func (repo *PostgresUserRepository) GetAll(ctx context.Context) ([]users.Domain, error) {
	var data []Users
	err := repo.ConnPostgres.Find(&data)
	if err.Error != nil {
		return []users.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (repo *PostgresUserRepository) Login(ctx context.Context, login requests.UserLogin) (users.Domain, error) {
	var user Users
	err := repo.ConnPostgres.Find(&user, "email = ? ", login.Email)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	check := utilities.CheckPassword(login.Password, user.Password)
	if check {
		return user.ToDomain(), nil
	}
	return users.Domain{}, nil
}
