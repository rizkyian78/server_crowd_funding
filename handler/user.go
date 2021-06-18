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

func (h *userHandler) LoginUser(c *gin.Context) {
	var input Users.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"error": errors}
		formatResponse := helper.APIResponse("failed to login", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, formatResponse)
		return
	}
	loginUser, err := h.userService.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		formatResponse := helper.APIResponse("failed to login", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, formatResponse)
		return
	}
	formatter := Users.FormatUser(loginUser, "token")
	formatResponse := helper.APIResponse("login Success", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, formatResponse)
}
