package controller

import (
	"basket-discount-api/internal/discount"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type controller struct {
	service discount.Service
}

func NewController(service discount.Service) *controller {
	return &controller{
		service: service,
	}
}

// MakeHandler returns a handler for the discount service.
func MakeDiscountHandler(instance *echo.Echo, s *controller) {
	instance.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "Service is up")
	})
	g := instance.Group("")
	baseUlr := "api/v1/discount"

	g.GET(fmt.Sprintf("%s/campaign/:campaignId", baseUlr), s.getCampaign)
}

// getCampaign godoc
// @Summary Get campaign
// @Tags discount
// @Param campaignId path string true "campaignId valid"
// @Success 200 {object} GetCampaignResponse
// @Success 404 {object} object
// @Failure 500 {string} string
// @Router /api/v1/discount/campaign/{campaignId} [get]
func (receiver *controller) getCampaign(c echo.Context) error {
	strCampaignId := c.Param("campaignId")
	if strCampaignId == "" {
		return c.JSON(http.StatusBadRequest, "CampaignId is required")
	}
	campaignId, _ := strconv.ParseInt(strCampaignId, 10, 0)
	var result = GetCampaignResponse{}
	data, err := receiver.service.GetCampaign(campaignId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if data == nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	result.Item = *data

	return c.JSON(http.StatusOK, result)
}
