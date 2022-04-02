package models

func (entity Campaign) ToDto() CampaignDto {
	var campaignDto = CampaignDto{
		CampaignId:   entity.Id,
		CampaignName: entity.Name,
	}
	return campaignDto
}
