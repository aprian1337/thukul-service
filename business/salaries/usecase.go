package salaries

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/helpers"
	"context"
	"time"
)

type SalaryUsecase struct {
	Repo    Repository
	Timeout time.Duration
}

func NewSalaryUsecase(repo Repository, timeout time.Duration) *SalaryUsecase {
	return &SalaryUsecase{
		Repo:    repo,
		Timeout: timeout,
	}
}

func (su *SalaryUsecase) GetList(ctx context.Context, search string) ([]Domain, error) {
	salaries, err := su.Repo.GetList(ctx, search)
	if err != nil {
		return []Domain{}, err
	}
	return salaries, nil
}
func (su *SalaryUsecase) GetById(ctx context.Context, id uint) (Domain, error) {
	salaries, err := su.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return salaries, nil
}
func (su *SalaryUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	salaries, err := su.Repo.Create(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return salaries, nil
}
func (su *SalaryUsecase) Update(ctx context.Context, domain Domain) (Domain, error) {
	if helpers.IsZero(domain.Minimal) || helpers.IsZero(domain.Maximal) {
		return Domain{}, businesses.ErrBadRequest
	}
	salaries, err := su.Repo.Update(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return salaries, nil
}

func (su *SalaryUsecase) Delete(ctx context.Context, id uint) error {
	err := su.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
