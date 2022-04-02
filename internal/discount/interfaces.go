package discount

import "basket-discount-api/internal/discount/models"

type ICampaignRepository interface {
	Create(entity models.Campaign) error
	DeActive(campaignId int64) error
	Update(entity models.Campaign) error
	Get(campaignId int64) (*models.Campaign, error)
	GetMany(isActive bool) ([]models.Campaign, error)
}

type IPromoCodeRepository interface {
	Create(entity PromoCode) error
	DeActive(promoCodeId int64) error
	Get(promoCodeId int64)
}

type Service interface {
	GetCampaign(campaignId int64) (*models.CampaignDto, error)
	CreateCampaign() error
	UpdateCampaign() error
	DeActiveCampaign() error
	CalculateCampaigns(cartDto models.CartDto) error
}
