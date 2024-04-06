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

func (h *compHanders) GetEmail(c *gin.Context) {
	email := c.PostForm("email")

	err := h.service.EmailSend(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Error: err.Error()})
		return
	}

	response := dto.Response{Message: "Email sucessfully sent!"}
	c.JSON(http.StatusOK, response)
}