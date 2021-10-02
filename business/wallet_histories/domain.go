package wallet_histories

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Domain struct {
	ID            int
	WalletId      int
	TransactionId *uuid.UUID
	Type          string
	Nominal       float64
	CreatedAt     time.Time
}

type Usecase interface {
	WalletHistoriesCreate(ctx context.Context, domain Domain) error
}

type Repository interface {
	WalletHistoriesCreate(ctx context.Context, domain Domain) error
}
