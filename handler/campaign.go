package handler

import (
	"bwa-golang/campaign"
	"bwa-golang/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	data, err := h.service.GetCampaigns(userId)

	if err != nil {
		response := helpers.ApiResponseBadRequest("Error to get campaigns", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponseSuccess("List of campaigns", campaign.FormatCampaigns(data))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {

	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helpers.ApiResponseBadRequest("Failed to get detail of campaign", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignById(input)
	if err != nil {
		response := helpers.ApiResponseBadRequest("Failed to get detail of campaign", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponseSuccess("Failed to get detail of campaign", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}
