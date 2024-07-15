package transaction

import "bwa-golang/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CrateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignId int `json:"campaign_id" binding:"required"`
	User       user.User
}
