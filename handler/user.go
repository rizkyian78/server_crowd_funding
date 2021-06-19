package handler

import (
	"crowd_fund_server/Users"
	"crowd_fund_server/auth"
	"crowd_fund_server/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService Users.Service
	authService auth.Service
}

func NewUserHandler(userService Users.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		formatResponse := helper.APIResponse("register account fail", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, formatResponse)
		return
	}
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		formatResponse := helper.APIResponse("register account fail", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, formatResponse)
		return
	}
	formatResponse := Users.FormatUser(newUser, token)
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

	token, err := h.authService.GenerateToken(loginUser.ID)
	if err != nil {
		formatResponse := helper.APIResponse("register account fail", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, formatResponse)
		return
	}
	formatter := Users.FormatUser(loginUser, token)
	formatResponse := helper.APIResponse("login Success", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, formatResponse)
}

func (h *userHandler) CheckEmail(c *gin.Context) {
	var input Users.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"error": errors}
		formatResponse := helper.APIResponse("Check Email Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, formatResponse)
		return
	}
	isEmailAvailable, err := h.userService.CheckEmail(input)
	if err != nil {
		errorMessage := gin.H{"error": "Server Error"}
		formatResponse := helper.APIResponse("Check Email Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, formatResponse)
		return
	}
	data := gin.H{
		"is_available": isEmailAvailable,
	}
	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email Available"
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to Upload Avatar Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(Users.User)
	userID := currentUser.ID
	path := fmt.Sprintf("public/%s-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": fmt.Sprintf("File path %s Not Found", path)}
		response := helper.APIResponse("Failed to Upload Avatar Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to Upload Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Successfully upload avatar", http.StatusOK, "error", data)
	c.JSON(http.StatusOK, response)
}
