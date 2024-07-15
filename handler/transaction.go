package handler

import (
	"bwa-golang/helpers"
	"bwa-golang/transaction"
	"bwa-golang/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	fmt.Println("currentUser", currentUser.ID)

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helpers.APIResponseUnprocessableEntity("Failed to get campaign transactions", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transactions, err := h.service.GetTransactionsByCampaignId(input)
	if err != nil {
		response := helpers.APIResponseUnprocessableEntity("Failed to get campaign transactions", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.APIResponseSuccess("Success to get campaign transactions", transaction.FormatCampaignTransactionList(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserID(userID)
	if err != nil {
		response := helpers.ApiResponseBadRequest("Failed to get user transactions", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponseSuccess("Success to get user transactions", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CrateTransactionInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helpers.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}

		respons := helpers.APIResponseUnprocessableEntity("Failed to create transaction", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, respons)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaction, err := h.service.CreateTransactions(input)

	if err != nil {
		response := helpers.ApiResponseBadRequest("Failed to create transaction", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponseCreated("Success to create transaction", newTransaction)
	c.JSON(http.StatusOK, response)

}
