package transaction

import (
	"bwa-golang/campaign"
	"bwa-golang/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     string
	Amount     int
	Status     string
	Code       string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
