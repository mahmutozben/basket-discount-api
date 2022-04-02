package models

import "time"

type (
	Campaign struct {
		Id                     int                     `gorm:"column:id;primary_key"`
		Name                   string                  `gorm:"column:name"`
		CreatedBy              string                  `gorm:"column:createdBy"`
		CreatedDateTime        *time.Time              `gorm:"column:createdDateTime"`
		StartDateTime          time.Time               `gorm:"column:startDateTime"`
		EndDateTime            time.Time               `gorm:"column:endDateTime"`
		IsActive               bool                    `gorm:"column:isActive"`
		CampaignDiscountDetail *CampaignDiscountDetail `gorm:"embedded"`
	}

	CampaignDiscountDetail struct {
		CampaignId   int          `gorm:"column:CampaignId"`
		Discount     float64      `gorm:"column:Discount"`
		DiscountType DiscountType `gorm:"column:DiscountType"`
	}

	DiscountType uint

	PromoCode struct {
		Id                int                `gorm:"column:id;primary_key"`
		PromoCode         string             `gorm:"column:promoCode"`
		Description       string             `gorm:"column:description"`
		IsActive          bool               `gorm:"column:isActive"`
		StartDateTime     *time.Time         `gorm:"column:startDateTime"`
		EndDateTime       *time.Time         `gorm:"column:endDateTime"`
		CreatedBy         string             `gorm:"column:createdBy"`
		CreatedDateTime   time.Time          `gorm:"column:createdDateTime"`
		ModifiedDateTime  *time.Time         `gorm:"column:modifiedDateTime"`
		PromoCodeCampaign *PromoCodeCampaign `gorm:"embedded"`
	}

	PromoCodeCampaign struct {
		PromoCodeId  int          `gorm:"column:PromoCodeId"`
		CampaignId   int64        `gorm:"column:CampaignId"`
		Discount     float64      `gorm:"column:Discount"`
		DiscountType DiscountType `gorm:"column:DiscountType"`
	}
)

const (
	DiscountRate DiscountType = iota
	DiscountAmount
)
