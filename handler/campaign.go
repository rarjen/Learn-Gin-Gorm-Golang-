package handler

import (
	"bwa-golang/campaign"
	"bwa-golang/helpers"
	"bwa-golang/user"
	"fmt"
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

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helpers.APIResponseUnprocessableEntity("Failed to update campaign1", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	checkDataExist, err := h.service.GetCampaignById(input)
	if err != nil {
		response := helpers.APIResponseUnprocessableEntity("Failed to update campaign1", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if checkDataExist.ID == 0 {
		response := helpers.APIResponseNotFound("Not found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	var inputData campaign.UpdateCampaignInput

	currentUser := c.MustGet("currentUser").(user.User)

	inputData.User = currentUser

	err = c.ShouldBindBodyWithJSON(&inputData)
	if err != nil {
		response := helpers.APIResponseUnprocessableEntity("Failed to update campaign2", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.service.UpdateCampaign(input, inputData)

	if err != nil {

		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponseUnprocessableEntity("Failed to update campaign3", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	afterUpdated, err := h.service.GetCampaignById(input)
	if err != nil {
		response := helpers.APIResponseUnprocessableEntity("Failed to update campaign4", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.APIResponseSuccess("Success to update campaign", campaign.FormatCampign(afterUpdated))

	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CreateCampignImageInput

	err := c.ShouldBind(&input)
	if err != nil {

		errors := helpers.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}
		response := helpers.APIResponseUnprocessableEntity("Failed to upload campaign image", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.APIResponseUnprocessableEntity("Failed to upload campaign image", data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.APIResponseUnprocessableEntity("Failed to upload campaign image", data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {

		data := gin.H{"is_uploaded": false}
		response := helpers.APIResponseUnprocessableEntity("Failed to upload campaign image", data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helpers.APIResponseSuccess("Success to upload campaign image", data)

	c.JSON(http.StatusOK, response)
}
