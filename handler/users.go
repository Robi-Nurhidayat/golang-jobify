package handler

import (
	"github.com/gin-gonic/gin"
	"jobify/helper"
	"jobify/user"
	"net/http"
)

type userHandler struct {
	service user.UserService
}

func NewUserHandler(service user.UserService) *userHandler {
	return &userHandler{
		service: service,
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

	response := helper.ApiResponse("Success register", http.StatusOK, "success", user.FormatterUser(userCreated))
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
		response := helper.ApiResponse("failed login", http.StatusBadRequest, "failed", helper.FormatValidationError(err))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success register", http.StatusOK, "success", user.FormatterUser(userLogin))
	c.JSON(http.StatusBadRequest, response)

}
