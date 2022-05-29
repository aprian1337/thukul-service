package graphql

import (
	"aprian1337/thukul-service/app/middlewares"
	"aprian1337/thukul-service/business/payments"
	"aprian1337/thukul-service/deliveries/payments/requests"
	"aprian1337/thukul-service/deliveries/payments/responses"
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
)

type Resolver struct{
	paymentUsecase payments.PaymentUsecase
}

func (ctrl *Resolver) TopUpCoin(params graphql.ResolveParams) (interface{}, error) {
	var(
		userId int
		nominal float64
		ok bool
	)
	header := params.Info.RootValue.(map[string]interface{})
	if header["token"]==""{
		return nil, fmt.Errorf("you're not have authorized for this action")
	}
	_, err := middlewares.GetClaimsAdminId(header["token"].(string))
	if err != nil {
		return nil, err
	}
	if userId, ok = params.Args["userId"].(int); !ok || userId == 0 {
		return nil, fmt.Errorf("userId is required and not zero value")
	}
	if nominal, ok = params.Args["nominal"].(float64); !ok || nominal == 0 {
		return nil, fmt.Errorf("nominal is required and not zero value")
	}
	data := requests.PaymentRequest{
		UserId:  userId,
		Nominal: nominal,
	}
	ctx := context.Background()
	resp, err := ctrl.paymentUsecase.TopUp(ctx, data.ToDomain())
	res := responses.FromDomainWallets(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (ctrl *Resolver) BuyCoin(params graphql.ResolveParams) (interface{}, error) {
	var(
		symbol string
		qty float64
		ok bool
	)
	if symbol, ok = params.Args["symbol"].(string); !ok || symbol == "" {
		return nil, fmt.Errorf("symbol is required and not null value")
	}
	if qty, ok = params.Args["qty"].(float64); !ok || qty == 0 {
		return nil, fmt.Errorf("qty is required and not zero value")
	}
	header := params.Info.RootValue.(map[string]interface{})
	if header["token"]==""{
		return nil, fmt.Errorf("you're not have authorized for this action")
	}
	id, err := middlewares.GetClaimsUserId(header["token"].(string))
	if err != nil {
		return nil, err
	}
	data := requests.PaymentRequest{
		Coin:  symbol,
		Qty: qty,
		UserId: int(id),
	}
	ctx := context.Background()
	err = ctrl.paymentUsecase.BuyCoin(ctx, data.ToDomain())
	if err != nil {
		return nil, err
	}
	return responses.BuySaleResponse{
		Status:  "success",
		Message: "check your email for confirm the purchase",
	}, nil
}

func (ctrl *Resolver) SellCoin(params graphql.ResolveParams) (interface{}, error) {
	var(
		symbol string
		qty float64
		ok bool
	)
	if symbol, ok = params.Args["symbol"].(string); !ok || symbol == "" {
		return nil, fmt.Errorf("symbol is required and not null value")
	}
	if qty, ok = params.Args["qty"].(float64); !ok || qty == 0 {
		return nil, fmt.Errorf("qty is required and not zero value")
	}
	header := params.Info.RootValue.(map[string]interface{})
	if header["token"]==""{
		return nil, fmt.Errorf("you're not have authorized for this action")
	}
	id, err := middlewares.GetClaimsUserId(header["token"].(string))
	if err != nil {
		return nil, err
	}
	data := requests.PaymentRequest{
		Coin:  symbol,
		Qty: qty,
		UserId: int(id),
	}
	ctx := context.Background()
	err = ctrl.paymentUsecase.SellCoin(ctx, data.ToDomain())
	if err != nil {
		return nil, err
	}
	return responses.BuySaleResponse{
		Status:  "success",
		Message: "check your email for sales confirmation",
	}, nil
}

func NewPaymentsResolver(paymentUsecase payments.PaymentUsecase) *Resolver{
	return &Resolver{
		paymentUsecase: paymentUsecase,
	}
}