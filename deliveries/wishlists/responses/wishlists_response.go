package responses

import (
	"aprian1337/thukul-service/business/wishlists"
	"time"
)

type WishlistsResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	UserId      int       `json:"user_id"`
	Nominal     float64   `json:"nominal"`
	TargetDate  string    `json:"target_date"`
	Priority    string    `json:"priority"`
	Note        string    `json:"note"`
	IsDone      int       `json:"is_done"`
	PicUrl      string    `json:"pic_url"`
	WishlistUrl string    `json:"wishlist_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WishlistsDestroyResponse struct {
	Message string `json:"message"`
}

func FromDomain(domain wishlists.Domain) WishlistsResponse {
	return WishlistsResponse{
		ID:          domain.ID,
		Name:        domain.Name,
		UserId:      domain.UserId,
		Nominal:     domain.Nominal,
		TargetDate:  domain.TargetDate,
		Priority:    domain.Priority,
		Note:        domain.Note,
		IsDone:      domain.IsDone,
		PicUrl:      domain.PicUrl,
		WishlistUrl: domain.WishlistUrl,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromListDomain(domain []wishlists.Domain) []WishlistsResponse {
	var result []WishlistsResponse
	for _, v := range domain {
		result = append(result, FromDomain(v))
	}
	return result
}
