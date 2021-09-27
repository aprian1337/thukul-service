package wishlists

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/wishlists"
	"aprian1337/thukul-service/repository/databases/users"
	"context"
	"gorm.io/gorm"
)

type PostgresWishlistRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresWishlistRepository(conn *gorm.DB) *PostgresWishlistRepository {
	return &PostgresWishlistRepository{
		ConnPostgres: conn,
	}
}

func (repo *PostgresWishlistRepository) GetList(ctx context.Context, userId int) ([]wishlists.Domain, error) {
	var data []Wishlists
	err := repo.ConnPostgres.Find(&data, "user_id=?", userId)
	if err.Error != nil {
		return []wishlists.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (repo *PostgresWishlistRepository) GetById(ctx context.Context, userId int, wishlistId int) (wishlists.Domain, error) {
	var data Wishlists
	err := repo.ConnPostgres.First(&data, "user_id=? AND id=?", userId, wishlistId)
	if err.Error != nil {
		return wishlists.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (repo *PostgresWishlistRepository) Create(ctx context.Context, domain wishlists.Domain, userId int) (wishlists.Domain, error) {
	pocket := FromDomain(domain)
	var user users.Users
	err := repo.ConnPostgres.First(&user, "user_id=?", userId)
	if err.Error != nil {
		return wishlists.Domain{}, businesses.ErrUserIdNotFound
	}
	err = repo.ConnPostgres.Create(&pocket)
	if err.Error != nil {
		return wishlists.Domain{}, err.Error
	}
	return pocket.ToDomain(), nil
}

func (repo *PostgresWishlistRepository) Update(ctx context.Context, domain wishlists.Domain, userId int, wishlistId int) (wishlists.Domain, error) {
	data := FromDomain(domain)
	dataTemp := FromDomain(domain)
	err := repo.ConnPostgres.First(&data, "user_id = ? AND id = ?", userId, wishlistId)
	if err.Error != nil {
		return wishlists.Domain{}, err.Error
	}
	data.Name = dataTemp.Name
	data.Nominal = dataTemp.Nominal
	data.TargetDate = dataTemp.TargetDate
	data.Priority = dataTemp.Priority
	data.Note = dataTemp.Note
	data.IsDone = dataTemp.IsDone
	data.PicUrl = dataTemp.PicUrl
	data.WishlistUrl = dataTemp.WishlistUrl
	repo.ConnPostgres.Save(&data)
	return data.ToDomain(), nil
}

func (repo *PostgresWishlistRepository) Delete(ctx context.Context, userId int, wishlistId int) (int64, error) {
	data := Wishlists{}
	err := repo.ConnPostgres.Delete(&data, "user_id=? AND id=?", userId, wishlistId)
	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
