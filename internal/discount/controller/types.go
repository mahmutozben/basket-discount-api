package controller

import "basket-discount-api/internal/discount/models"

type (
	GetCampaignResponse struct {
		Item models.CampaignDto `json:"item"`
	}
)
