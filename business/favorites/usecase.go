package favorites

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/coins"
	"aprian1337/thukul-service/business/users"
	"context"
	"time"
)

type FavoriteUsecase struct {
	Repo        Repository
	CoinUsecase coins.Usecase
	UserUsecase users.Usecase
	Timeout     time.Duration
}

func NewFavoriteUsecase(repo Repository, coin coins.Usecase, timeout time.Duration) *FavoriteUsecase {
	return &FavoriteUsecase{
		Repo:        repo,
		CoinUsecase: coin,
		Timeout:     timeout,
	}
}

func (uc *FavoriteUsecase) GetList(ctx context.Context, userId int) ([]Domain, error) {
	data, err := uc.Repo.GetList(ctx, userId)
	if err != nil {
		return []Domain{}, err
	}
	return data, nil
}

func (uc *FavoriteUsecase) GetById(ctx context.Context, userId int, favoriteId int) (Domain, error) {
	data, err := uc.Repo.GetById(ctx, userId, favoriteId)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}

func (uc *FavoriteUsecase) Create(ctx context.Context, domain Domain, userId int) (Domain, error) {
	symbol, err := uc.CoinUsecase.GetBySymbol(ctx, domain.CoinSymbol)
	if err != nil {
		return Domain{}, businesses.ErrUserIdNotFound
	}
	_, err = uc.UserUsecase.GetById(ctx, userId)
	if err != nil {
		return Domain{}, businesses.ErrUserIdNotFound
	}
	domain.CoinId = symbol.Id
	data, err := uc.Repo.Create(ctx, domain, userId)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}

func (uc *FavoriteUsecase) Delete(ctx context.Context, userId int, favoriteId int) error {
	rowsAffected, err := uc.Repo.Delete(ctx, userId, favoriteId)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return businesses.ErrNothingDestroy
	}
	return nil
}
