package salaries

import (
	"time"
)

type Response struct {
	ID        uint      `json:"id"`
	Minimal   float64   `json:"minimal"`
	Maximal   float64   `json:"maximal"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
