package models

import "github.com/google/uuid"

type (
	CartDto struct {
		BasketId   uuid.UUID
		CustomerId uuid.UUID
		Items      []CartItemDto
	}
	CartItemDto struct {
		ProductCode string
		Price       float64
		Quantity    int
		MerchantId  uuid.UUID
	}

	CampaignDto struct {
		CampaignId   int64  `json:"campaignId"`
		CampaignName string `json:"campaignName"`
	}
)
