package handler

import (
	"Ozinshe/pkg/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var user repository.SignupUser
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "невалидный JSON"})
		return
	}

	log.Printf("email: %s, Полученные пароль: %s, Подтверждение пароля: %s", user.Email, user.Password, user.ConfirmPassword)

	err := validateUserData(user, h.db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(user.Email, user.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    id,
		"token": token,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input repository.SignInInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "невалидный JSON"})
		return
	}

	log.Printf("email: %s, Полученные пароль: %s", input.Email, input.Password)

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
