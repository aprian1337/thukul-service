package users

import (
	"aprian1337/thukul-service/business/users"
	"context"
	"errors"
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

func (repo *PostgresUserRepository) Update(ctx context.Context, domain *users.Domain) (users.Domain, error) {
	data := FromDomain(domain)
	if repo.ConnPostgres.Save(&data).Error != nil {
		return users.Domain{}, errors.New("bad requests")
	}
	return data.ToDomain(), nil
}

func (repo *PostgresUserRepository) Delete(ctx context.Context, id uint) error {
	user := Users{}
	err := repo.ConnPostgres.Delete(&user, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
