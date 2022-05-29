package graphql

import (
	"aprian1337/thukul-service/business/salaries"
	"context"
	"github.com/graphql-go/graphql"
)

type Resolver struct{
	salaryUsecase salaries.SalaryUsecase
}

func (ctrl *Resolver) GetSalaries(params graphql.ResolveParams) (interface{}, error) {
	ctx := context.Background()
	salary, err := ctrl.salaryUsecase.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return &salary, nil
}

func NewSalaryResolver(salaryUsecase salaries.SalaryUsecase) *Resolver{
	return &Resolver{
		salaryUsecase: salaryUsecase,
	}
}