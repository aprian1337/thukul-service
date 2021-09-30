package transactions

import (
	"aprian1337/thukul-service/repository/databases/coins"
	"aprian1337/thukul-service/repository/databases/users"
	"github.com/google/uuid"
	"time"
)

type Transactions struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserId            int
	User              users.Users `gorm:"foreignKey:user_id"`
	CoinId            int
	Coin              coins.Coins `gorm:"foreignKey:coin_id"`
	Qty               float64
	Status            int `gorm:"size:1"`
	DatetimeRequest   time.Time
	DatetimeVerify    time.Time
	DatetimeCompleted time.Time
}
