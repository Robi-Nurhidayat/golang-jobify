package handler

import (
	"github.com/gin-gonic/gin"
	"jobify/auth"
	"jobify/helper"
	"jobify/user"
	"net/http"
)

type userHandler struct {
	service     user.UserService
	authService auth.ServiceAuth
}

func NewUserHandler(service user.UserService, authService auth.ServiceAuth) *userHandler {
	return &userHandler{
		service:     service,
		authService: authService,
	}
}

func (h *userHandler) Register(c *gin.Context) {

	var input user.RegisterInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.ApiResponse("Please isi semua field", http.StatusBadRequest, "failed", helper.FormatValidationError(err))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userCreated, err := h.service.Register(input)
	if err != nil {
		response := helper.ApiResponse("Failed to register", http.StatusBadRequest, "failed", helper.FormatValidationError(err))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(userCreated.Id)
	if err != nil {
		response := helper.ApiResponse("Failed to register", http.StatusBadRequest, "failed", helper.FormatValidationError(err))
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success register", http.StatusOK, "success", user.FormatterUser(userCreated, token))
	c.JSON(http.StatusBadRequest, response)
}

func (h *userHandler) Login(c *gin.Context) {

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.ApiResponse("Please isi semua field", http.StatusBadRequest, "failed", helper.FormatValidationError(err))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userLogin, err := h.service.Login(input)
	if err != nil {
		response := helper.ApiResponse("email or password wrong !", http.StatusBadRequest, "failed", "not found that email")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(userLogin.Id)
	if err != nil {
		response := helper.ApiResponse("failed login", http.StatusBadRequest, "failed", helper.FormatValidationError(err))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success login", http.StatusOK, "success", user.FormatterUser(userLogin, token))
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) AllUsers(c *gin.Context) {

	users, err := h.service.GetAllUser()

	if err != nil {
		response := helper.ApiResponse("Failed get all users", http.StatusBadRequest, "failed", []user.User{})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if len(users) == 0 {
		response := helper.ApiResponse("Failed get all users", http.StatusBadRequest, "failed", []user.User{})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success get all users", http.StatusOK, "success", users)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) FetchUser(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)

	response := helper.ApiResponse("success get user", http.StatusOK, "success", user.FormatterUser(currentUser, ""))
	c.JSON(http.StatusOK, response)

}
