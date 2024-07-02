package handler

import (
	"bwa-golang/helpers"
	"bwa-golang/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
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

	result, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helpers.ApiResponseBadRequest("Something went wrong!")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token := "tokentest"

	formatter := user.FormatUser(result, token)

	response := helpers.APIResponseCreated("Success register!", formatter)

	c.JSON(http.StatusCreated, response)
}
