package campaign

import (
	"crowd_fund_server/Users"
	"time"
)

type Campaign struct {
	ID             string
	UserID         string
	Name           string
	ShortDesc      string
	Description    string
	Perks          string
	BackerCount    int
	GoalAmount     int
	CurrentAmount  int
	Slug           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CampaignImages []CampaignImages
	User           Users.User
}

type CampaignImages struct {
	ID         string
	CampaignID string
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
