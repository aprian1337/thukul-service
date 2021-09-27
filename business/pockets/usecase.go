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
	Repo    Repository
	RepoAct activities.Repository
	Timeout time.Duration
}

func NewPocketUsecase(repo Repository, repoActivity activities.Repository, timeout time.Duration) *PocketUsecase {
	return &PocketUsecase{
		Repo:    repo,
		RepoAct: repoActivity,
		Timeout: timeout,
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
func (pu *PocketUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	pockets, err := pu.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return pockets, nil
}
func (pu *PocketUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	if domain.UserId == 0 || domain.Name == "" {
		return Domain{}, businesses.ErrBadRequest
	}
	pockets, err := pu.Repo.Create(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return pockets, nil
}
func (pu *PocketUsecase) Update(ctx context.Context, id int, domain Domain) (Domain, error) {
	domain.ID = id
	pockets, err := pu.Repo.Update(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return pockets, nil
}

func (pu *PocketUsecase) Delete(ctx context.Context, id int) error {
	rowsAffected, err := pu.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return businesses.ErrNothingDestroy
	}
	return nil
}

func (pu *PocketUsecase) GetTotal(ctx context.Context, id int, kind string) (int64, error) {
	if kind != "income" && kind != "expense" && kind != "profit" {
		return 0, businesses.ErrBadRequest
	}
	total, err := pu.RepoAct.GetTotal(ctx, id, kind)
	if err != nil {
		return 0, err
	}
	return total, nil
}
