package transaction

import (
	"time"
)

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

// Format 1 transaction
func FormatCampaignTransaction(transactions Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transactions.ID
	formatter.Name = transactions.User.Name
	formatter.Amount = transactions.Amount
	formatter.CreatedAt = transactions.CreatedAt
	return formatter
}

func FormatCampaignTransactionList(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var transactionsFormatter []CampaignTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
	}

	campaignFormatter := CampaignFormatter{
		Name: transaction.Campaign.Name,
		ImageURL: func() string {
			if len(transaction.Campaign.CampaignImages) > 0 && transaction.Campaign.CampaignImages[0].FileName != "" {
				return transaction.Campaign.CampaignImages[0].FileName
			}
			return ""
		}(),
	}

	formatter.Campaign = campaignFormatter

	return formatter

}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionFormatter []UserTransactionFormatter

	for _, v := range transactions {
		formatter := FormatUserTransaction(v)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}
