package handler

import (
	"crowd_fund_server/Users"
	"crowd_fund_server/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService Users.Service
}

func NewUserHandler(userService Users.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//catch user input
	var input Users.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"error": errors}
		formatResponse := helper.APIResponse("register account fail", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, formatResponse)
		return
	}
	user, err := h.userService.RegisterUser(input)
	if err != nil {
		formatResponse := helper.APIResponse("register account fail", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, formatResponse)
		return
	}
	// token, err := h.
	formatResponse := Users.FormatUser(user, "toketok")
	res := helper.APIResponse("Account Has been created", http.StatusOK, "success", formatResponse)
	c.JSON(http.StatusOK, res)
}
