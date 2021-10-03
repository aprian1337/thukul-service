package responses

type ConfirmResponse struct {
	Message      string  `json:"message"`
	PaymentTotal float64 `json:"payment_total"`
	PaymentType  string  `json:"payment_type"`
	WalletSaldo  float64 `json:"wallet_saldo"`
}
