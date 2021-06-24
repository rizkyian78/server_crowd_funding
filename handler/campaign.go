package handler

import (
	"crowd_fund_server/campaign"
	"crowd_fund_server/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID := c.Query("user_id")
	campaigns, err := h.service.FindCampaigns(userID)
	if err != nil {
		formatResponse := helper.APIResponse("Get Campaigns error", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, formatResponse)
		return
	}
	formatResponse := helper.APIResponse("register account fail", http.StatusBadRequest, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusBadRequest, formatResponse)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.FindCampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		formatResponse := helper.APIResponse("Get Campaigns error", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, formatResponse)
		return
	}
	campaignDetail, err := h.service.FindCampaignByID(input)
	if err != nil {
		formatResponse := helper.APIResponse("Get Campaigns error", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, formatResponse)
		return
	}
	fmt.Println(campaignDetail.ShortDesc, "<<<<<<<<<<<<<<<<<,")
	response := helper.APIResponse("Campaign Detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}
