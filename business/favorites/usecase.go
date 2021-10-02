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

func NewFavoriteUsecase(repo Repository, user users.Usecase, coin coins.Usecase, timeout time.Duration) *FavoriteUsecase {
	return &FavoriteUsecase{
		Repo:        repo,
		CoinUsecase: coin,
		UserUsecase: user,
		Timeout:     timeout,
	}
}

func (uc *FavoriteUsecase) GetList(ctx context.Context, userId int) ([]Domain, error) {
	data, err := uc.Repo.FavoritesGetList(ctx, userId)
	if err != nil {
		return []Domain{}, err
	}
	return data, nil
}

func (uc *FavoriteUsecase) GetById(ctx context.Context, userId int, favoriteId int) (Domain, error) {
	data, err := uc.Repo.FavoritesGetById(ctx, userId, favoriteId)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}

func (uc *FavoriteUsecase) Create(ctx context.Context, domain Domain, userId int) (Domain, error) {
	getSymbol, err := uc.CoinUsecase.GetBySymbol(ctx, domain.Symbol)
	if err != nil {
		return Domain{}, businesses.ErrTokenNotFound
	}
	_, err = uc.UserUsecase.GetById(ctx, userId)
	if err != nil {
		return Domain{}, businesses.ErrUserIdNotFound
	}
	res, err := uc.Repo.FavoritesCheck(ctx, userId, getSymbol.Id)
	if err != nil {
		return Domain{}, businesses.ErrUserIdNotFound
	}
	if res > 0 {
		return Domain{}, businesses.ErrFavoriteIsAlready
	}

	domain.CoinId = getSymbol.Id
	domain.UserId = userId
	data, err := uc.Repo.FavoritesCreate(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return data.AddCoins(getSymbol), nil
}

func (uc *FavoriteUsecase) Delete(ctx context.Context, userId int, favoriteId int) error {
	rowsAffected, err := uc.Repo.FavoritesDelete(ctx, userId, favoriteId)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return businesses.ErrNothingDestroy
	}
	return nil
}
