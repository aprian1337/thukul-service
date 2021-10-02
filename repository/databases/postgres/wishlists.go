package postgres

import (
	"aprian1337/thukul-service/business/wishlists"
	wishlists2 "aprian1337/thukul-service/repository/databases/records"
	"context"
	"gorm.io/gorm"
)

type WishlistRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresWishlistRepository(conn *gorm.DB) *WishlistRepository {
	return &WishlistRepository{
		ConnPostgres: conn,
	}
}

func (repo *WishlistRepository) WishlistsGetList(ctx context.Context, userId int) ([]wishlists.Domain, error) {
	var data []wishlists2.Wishlists
	err := repo.ConnPostgres.Find(&data, "user_id=?", userId)
	if err.Error != nil {
		return []wishlists.Domain{}, err.Error
	}
	return wishlists2.WishlistsToListDomain(data), nil
}

func (repo *WishlistRepository) WishlistsGetById(ctx context.Context, userId int, wishlistId int) (wishlists.Domain, error) {
	var data wishlists2.Wishlists
	err := repo.ConnPostgres.First(&data, "user_id=? AND id=?", userId, wishlistId)
	if err.Error != nil {
		return wishlists.Domain{}, err.Error
	}
	return data.WishlistsToDomain(), nil
}

func (repo *WishlistRepository) WishlistsCreate(ctx context.Context, domain wishlists.Domain, userId int) (wishlists.Domain, error) {
	pocket := wishlists2.WishlistsFromDomain(domain)
	err := repo.ConnPostgres.Create(&pocket)
	if err.Error != nil {
		return wishlists.Domain{}, err.Error
	}
	return pocket.WishlistsToDomain(), nil
}

func (repo *WishlistRepository) WishlistsUpdate(ctx context.Context, domain wishlists.Domain, userId int, wishlistId int) (wishlists.Domain, error) {
	data := wishlists2.WishlistsFromDomain(domain)
	dataTemp := wishlists2.WishlistsFromDomain(domain)
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
	return data.WishlistsToDomain(), nil
}

func (repo *WishlistRepository) WishlistsDelete(ctx context.Context, userId int, wishlistId int) (int64, error) {
	data := wishlists2.Wishlists{}
	err := repo.ConnPostgres.Delete(&data, "user_id=? AND id=?", userId, wishlistId)
	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
