package requests

import "aprian1337/thukul-service/business/payments"

type PaymentRequest struct {
	UserId  int     `json:"user_id"`
	Kind    string  `json:"kind"`
	Coin    string  `json:"coin"`
	Qty     float64 `json:"qty"`
	Nominal float64 `json:"nominal"`
}

func (pay *PaymentRequest) ToDomain() payments.Domain {
	return payments.Domain{
		UserId:  pay.UserId,
		Kind:    pay.Kind,
		Coin:    pay.Coin,
		Qty:     pay.Qty,
		Nominal: pay.Nominal,
	}
}
