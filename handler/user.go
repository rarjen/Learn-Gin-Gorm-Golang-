package handler

import (
	"bwa-golang/auth"
	"bwa-golang/helpers"
	"bwa-golang/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	//tangkap dari input user
	//map input dari user ke struct RegisterUserInput
	//struct diatas di passing ke service
	//service, call RegisterUser service
	//service mengecek apakah inputnya valid

	var input user.RegisterUserInput
	// Melakukan
	// Jika ada error validasi maka akan ditangkap disini
	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helpers.FormatErrorValidation(err)

		//gin.H{} adalah map untuk menampung key dan value
		errorMessage := gin.H{"error": errors}

		response := helpers.APIResponseUnprocessableEntity("Check your input!", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userExist, err := h.userService.CheckUserExist(input.Email)

	if err != nil {
		response := helpers.ApiResponseBadRequest("Something went wrong!", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if userExist.ID != 0 {
		response := helpers.ApiResponseBadRequest("Email already exist!", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Query
	result, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helpers.ApiResponseBadRequest("Something went wrong!", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(result.ID)

	if err != nil {
		response := helpers.ApiResponseBadRequest("Something went wrong!", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(result, token)

	response := helpers.APIResponseCreated("Success register!", formatter)

	c.JSON(http.StatusCreated, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindBodyWithJSON(&input)

	if err != nil {
		errors := helpers.FormatErrorValidation(err)

		//gin.H{} adalah map untuk menampung key dan value
		errorMessage := gin.H{"error": errors}

		response := helpers.APIResponseUnprocessableEntity("Check your input!", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Query
	result, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}

		response := helpers.APIResponseNotFound("Email/Password not match!", errorMessage)
		c.JSON(http.StatusNotFound, response)
		return
	}

	token, err := h.authService.GenerateToken(result.ID)

	if err != nil {
		response := helpers.ApiResponseBadRequest("Login failed!", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(result, token)

	response := helpers.APIResponseSuccess("Success login!", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) ChekEmailAvailablity(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindBodyWithJSON(&input)

	if err != nil {
		errors := helpers.FormatErrorValidation(err)

		//gin.H{} adalah map untuk menampung key dan value
		errorMessage := gin.H{"error": errors}

		response := helpers.APIResponseUnprocessableEntity("Email checking failed!", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {

		errorMessage := gin.H{"error": "Server error!"}

		response := helpers.APIResponseUnprocessableEntity("Email checking failed!", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}
	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helpers.APIResponseSuccess(metaMessage, data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.ApiResponseBadRequest("Failed upload avatar", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.ApiResponseBadRequest("Failed upload avatar", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userId, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.ApiResponseBadRequest("Failed upload avatar", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helpers.APIResponseSuccess("Success upload avatar", data)

	c.JSON(http.StatusOK, response)
}
