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

	campaigns, err := h.service.GetCampaigns(userId)

	if err != nil {
		response := helpers.ApiResponseBadRequest("Error to get campaigns", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponseSuccess("List of campaigns", campaigns)
	c.JSON(http.StatusOK, response)

}
