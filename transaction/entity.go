package transaction

import (
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
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
