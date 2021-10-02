package postgres

import (
	"aprian1337/thukul-service/business/users"
	users2 "aprian1337/thukul-service/repository/databases/records"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		ConnPostgres: conn,
	}
}

func (repo *UserRepository) UsersGetById(ctx context.Context, id int) (users.Domain, error) {
	var user users2.Users
	err := repo.ConnPostgres.Find(&user, "id = ?", id)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.UsersToDomain(), nil
}

func (repo *UserRepository) UsersGetByEmail(ctx context.Context, email string) (users.Domain, error) {
	var user users2.Users
	err := repo.ConnPostgres.Find(&user, "email = ?", email)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.UsersToDomain(), nil
}

func (repo *UserRepository) UsersCreate(ctx context.Context, register *users.Domain) (users.Domain, error) {
	user := users2.Users{
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
	return user.UsersToDomain(), nil
}

func (repo *UserRepository) UsersGetAll(ctx context.Context) ([]users.Domain, error) {
	var data []users2.Users
	err := repo.ConnPostgres.Find(&data)
	if err.Error != nil {
		return []users.Domain{}, err.Error
	}
	return users2.UsersToListDomain(data), nil
}

func (repo *UserRepository) UsersGetByIdWithWallet(ctx context.Context, id int) (users.Domain, error) {
	var data users2.Users
	err := repo.ConnPostgres.Preload("Wallets").Find(&data)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return data.UsersToDomain(), nil
}

func (repo *UserRepository) UsersUpdate(ctx context.Context, domain *users.Domain) (users.Domain, error) {
	data := users2.UsersFromDomain(domain)
	if repo.ConnPostgres.Save(&data).Error != nil {
		return users.Domain{}, errors.New("bad requests")
	}
	return data.UsersToDomain(), nil
}

func (repo *UserRepository) UsersDelete(ctx context.Context, id uint) error {
	user := users2.Users{}
	err := repo.ConnPostgres.Delete(&user, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
