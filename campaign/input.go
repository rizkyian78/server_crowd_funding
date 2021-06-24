package campaign

import "os/user"

type FindCampaignDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name        string `json:"name"`
	ShortDesc   string `json:"short_description"`
	Description string `json:"description"`
	GoalAmount  int    `json:"goal_amount"`
	Perks       string `json:"perks"`
	User        user.User
}
