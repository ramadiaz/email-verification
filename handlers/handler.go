package handlers

import (
	"email-verification/dto"
	"email-verification/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type compHanders struct {
	service services.CompServices
}

func NewCompHandlers(s services.CompServices) *compHanders {
	return &compHanders{
		service: s,
	}
}

func (h *compHanders) GetToken(c *gin.Context) {
	email := c.PostForm("email")

	err := h.service.TokenSend(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Error: err.Error()})
		return
	}

	response := dto.Response{Message: "Email successfully sent!"}
	c.JSON(http.StatusOK, response)
}

func (h *compHanders) VerifyToken(c *gin.Context) {
	email := c.PostForm("email")
	token := c.PostForm("token")

	err := h.service.TokenVerify(email, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Error: err.Error()})
		return
	}

	response := dto.Response{Message: "Email successfully verified!"}
	c.JSON(http.StatusOK, response)

}

func (h *compHanders) RegistUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	err := h.service.RegistUser(email, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Error: err.Error()})
		return
	}

	response := dto.Response{Message: "Account registered successfully!"}
	c.JSON(http.StatusOK, response)
}