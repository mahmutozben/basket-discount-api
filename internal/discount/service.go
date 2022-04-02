package discount

import "basket-discount-api/internal/discount/models"

type service struct {
	campaignRepository ICampaignRepository
}

func NewService(campaignRepository ICampaignRepository) Service {
	if campaignRepository == nil {
		return nil
	}

	return &service{
		campaignRepository: campaignRepository,
	}
}

func (receiver *service) GetCampaign(campaignId int64) (*models.CampaignDto, error) {
	campaign, err := receiver.campaignRepository.Get(campaignId)
	if err != nil {
		return nil, err
	}
	dto := campaign.ToDto()
	return &dto, nil
}

func (receiver *service) CreateCampaign() error {
	return nil
}

func (receiver *service) UpdateCampaign() error {
	return nil
}

func (receiver *service) DeActiveCampaign() error {
	return nil
}

func (receiver *service) CalculateCampaigns(cartDto models.CartDto) error {
	return nil
}
