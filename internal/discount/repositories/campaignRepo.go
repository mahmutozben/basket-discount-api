package repositories

import (
	"basket-discount-api/internal/discount"
	"basket-discount-api/internal/discount/models"
	"gorm.io/gorm"
)

type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepo(db *gorm.DB) discount.ICampaignRepository {
	if db == nil {
		return nil
	}
	return &campaignRepository{
		db: db,
	}
}

func (repo *campaignRepository) Create(entity models.Campaign) error {

	return nil
}

func (repo *campaignRepository) DeActive(campaignId int64) error {
	return nil
}

func (repo *campaignRepository) Update(entity models.Campaign) error {
	return nil
}

func (repo *campaignRepository) Get(campaignId int64) (*models.Campaign, error) {
	var campaign models.Campaign
	err := repo.db.Table("campaign").Find(&campaign, campaignId).Error
	if err != nil {
		return nil, err
	}
	return &campaign, nil
}

func (repo *campaignRepository) GetMany(isActive bool) ([]models.Campaign, error) {
	var campaigns []models.Campaign
	return campaigns, nil
}
