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
