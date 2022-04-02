package models

import "time"

type (
	Campaign struct {
		Id                     int64                  `gorm:"column:Id;primary_key"`
		Name                   string                 `gorm:"column:Name"`
		CreatedBy              string                 `gorm:"column:CreatedBy"`
		StartDateTime          *time.Time             `gorm:"column:StartDateTime"`
		EndDateTime            *time.Time             `gorm:"column:EndDateTime"`
		IsActive               bool                   `gorm:"column:IsActive"`
		CampaignDiscountDetail CampaignDiscountDetail `gorm:"embedded"`
	}

	CampaignDiscountDetail struct {
		CampaignId   int64        `gorm:"column:CampaignId"`
		Discount     float64      `gorm:"column:Discount"`
		DiscountType DiscountType `gorm:"column:DiscountType"`
	}

	DiscountType uint

	PromoCode struct {
		Id                int64              `gorm:"column:Id;primary_key"`
		PromoCode         string             `gorm:"column:PromoCode"`
		Description       string             `gorm:"column:Description"`
		IsActive          bool               `gorm:"column:IsActive"`
		StartDateTime     *time.Time         `gorm:"column:StartDateTime"`
		EndDateTime       *time.Time         `gorm:"column:EndDateTime"`
		CreatedBy         string             `gorm:"column:CreatedBy"`
		CreatedDateTime   time.Time          `gorm:"column:CreatedDateTime"`
		ModifiedDateTime  *time.Time         `gorm:"column:ModifiedDateTime"`
		PromoCodeCampaign *PromoCodeCampaign `gorm:"embedded"`
	}

	PromoCodeCampaign struct {
		PromoCodeId  int64        `gorm:"column:PromoCodeId"`
		CampaignId   int64        `gorm:"column:CampaignId"`
		Discount     float64      `gorm:"column:Discount"`
		DiscountType DiscountType `gorm:"column:DiscountType"`
	}
)

const (
	DiscountRate DiscountType = iota
	DiscountAmount
)
