package campaign

import (
	"strings"
)

// cara restructuring object

type CampaignFormatter struct {
	ID            string `json:"id"`
	UserId        string `json:"user_id"`
	Name          string `json:"name"`
	ShortDesc     string `json:"short_description"`
	ImageUrl      string `json:"image_path"`
	GoalAmount    int    `json:"goal_amount"`
	CurrentAmount int    `json:"current_amount"`
	Slug          string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserId = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDesc = campaign.ShortDesc
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.ImageUrl = ""
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}
	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}
	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID            string                   `json:"id"`
	Name          string                   `json:"name"`
	ShortDesc     string                   `json:"short_desc"`
	Description   string                   `json:"description"`
	ImageUrl      string                   `json:"image_url"`
	GoalAmount    int                      `json:"goal_amount"`
	CurrentAmount int                      `json:"current_amount"`
	UserId        string                   `json:"user_id"`
	Slug          string                   `json:"slug"`
	Perks         []string                 `json:"perks"`
	User          CampaignUserFormatter    `json:"user"`
	Images        []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignFormatter := CampaignDetailFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserId = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDesc = campaign.ShortDesc
	campaignFormatter.Description = campaign.Description
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.ImageUrl = ""
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}
	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignFormatter.Perks = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageUrl = user.AvatarFileName
	campaignFormatter.User = campaignUserFormatter

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageUrl = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary
		images = append(images, campaignImageFormatter)
	}
	campaignFormatter.Images = images
	return campaignFormatter
}
