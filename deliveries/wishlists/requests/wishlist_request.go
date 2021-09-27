package requests

import (
	"aprian1337/thukul-service/business/wishlists"
)

type WishlistsRequest struct {
	Name        string  `json:"name"`
	UserId      int     `json:"user_id"`
	Nominal     float64 `json:"nominal"`
	TargetDate  string  `json:"target_date"`
	Priority    string  `json:"priority"`
	Note        string  `json:"note"`
	IsDone      int     `json:"is_done"`
	PicUrl      string  `json:"pic_url"`
	WishlistUrl string  `json:"wishlist_url"`
}

func (wr *WishlistsRequest) ToDomain() wishlists.Domain {
	return wishlists.Domain{
		Name:        wr.Name,
		UserId:      wr.UserId,
		Nominal:     wr.Nominal,
		TargetDate:  wr.TargetDate,
		Priority:    wr.Priority,
		Note:        wr.Note,
		IsDone:      wr.IsDone,
		PicUrl:      wr.PicUrl,
		WishlistUrl: wr.WishlistUrl,
	}
}
