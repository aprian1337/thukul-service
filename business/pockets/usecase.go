package pockets

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/activities"
	"aprian1337/thukul-service/helpers"
	"context"
	"strconv"
	"time"
)

type PocketUsecase struct {
	Repo            Repository
	ActivityUsecase activities.Usecase
	Timeout         time.Duration
}

func NewPocketUsecase(repo Repository, activityUsecase activities.Usecase, timeout time.Duration) *PocketUsecase {
	return &PocketUsecase{
		Repo:            repo,
		ActivityUsecase: activityUsecase,
		Timeout:         timeout,
	}
}

func (pu *PocketUsecase) GetList(ctx context.Context, id string) ([]Domain, error) {
	convId := 0
	if id != "" {
		if helpers.IsInt(id) == false {
			return []Domain{}, businesses.ErrInvalidId
		}
		convId, _ = strconv.Atoi(id)
	}

	pockets, err := pu.Repo.GetList(ctx, convId)
	if err != nil {
		return []Domain{}, err
	}
	return pockets, nil
}
func (pu *PocketUsecase) GetById(ctx context.Context, userId int, pocketId int) (Domain, error) {
	pockets, err := pu.Repo.GetById(ctx, userId, pocketId)
	if err != nil {
		return Domain{}, err
	}
	return pockets, nil
}
func (pu *PocketUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, businesses.ErrBadRequest
	}
	pockets, err := pu.Repo.Create(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return pockets, nil
}
func (pu *PocketUsecase) Update(ctx context.Context, domain Domain, userId int, pocketId int) (Domain, error) {
	pockets, err := pu.Repo.Update(ctx, domain, userId, pocketId)
	if err != nil {
		return Domain{}, err
	}
	return pockets, nil
}

func (pu *PocketUsecase) Delete(ctx context.Context, userId int, pocketId int) error {
	rowsAffected, err := pu.Repo.Delete(ctx, userId, pocketId)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return businesses.ErrNothingDestroy
	}
	return nil
}

func (pu *PocketUsecase) GetTotalByActivities(ctx context.Context, userId int, pocketId int, kind string) (int64, error) {
	if kind != "income" && kind != "expense" && kind != "profit" {
		return 0, businesses.ErrBadRequest
	}
	_, err := pu.GetById(ctx, userId, pocketId)
	if err != nil {
		return 0, businesses.ErrUserIdOrPocketNotFound
	}
	total, err := pu.ActivityUsecase.GetTotal(ctx, userId, pocketId, kind)
	if err != nil {
		return 0, err
	}
	return total, nil
}
