package records

import (
	"aprian1337/thukul-service/business/transactions"
	"github.com/google/uuid"
	"time"
)

type Transactions struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserId            int
	User              Users `gorm:"foreignKey:user_id"`
	CoinId            int
	Coin              Coins `gorm:"foreignKey:coin_id"`
	Qty               float64
	Price             float64
	Status            int `gorm:"size:1"`
	Type              string
	DatetimeRequest   time.Time
	DatetimeVerify    *time.Time
	DatetimeCompleted *time.Time
}

func TransactionsFromDomain(domain transactions.Domain) Transactions {
	return Transactions{
		UserId:          domain.UserId,
		CoinId:          domain.CoinId,
		Qty:             domain.Qty,
		Price:           domain.Price,
		Status:          0,
		Type:            domain.Kind,
		DatetimeRequest: time.Now(),
	}
}

func (t *Transactions) TransactionsToDomain() transactions.Domain {
	return transactions.Domain{
		Id:                t.ID.String(),
		UserId:            t.UserId,
		CoinId:            t.CoinId,
		Qty:               t.Qty,
		Kind:              t.Type,
		Price:             t.Price,
		Status:            t.Status,
		DatetimeRequest:   t.DatetimeRequest,
		DatetimeVerify:    t.DatetimeVerify,
		DatetimeCompleted: t.DatetimeCompleted,
	}
}
