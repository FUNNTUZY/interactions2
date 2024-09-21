package dto

import "time"

type InteractionDTO struct {
	ID              string    `bun:",pk" json:"id"`
	UserID          string    `json:"user_id"`
	AdID            string    `json:"ad_id"`
	SellerID        string    `json:"seller_id"`
	InteractionType string    `json:"type"`
	CreatedAt       time.Time `json:"created_at"`
}
