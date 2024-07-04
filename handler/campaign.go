package handler

import (
	"bwa-golang/campaign"
	"bwa-golang/helpers"
	"bwa-golang/user"
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

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponseUnprocessableEntity("Failed to create campaign", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCampaign, err := h.service.CreateCampign(input)

	if err != nil {
		errors := helpers.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponseUnprocessableEntity("Failed to create campaign", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.APIResponseCreated("Success to create campaign", campaign.FormatCampign(newCampaign))

	c.JSON(http.StatusCreated, response)
}
