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

	response := dto.Response{Message: "Email sucessfully sent!"}
	c.JSON(http.StatusOK, response)
}

func (h *compHanders) VerifyToken(c  *gin.Context) {
	email := c.PostForm("email")
	token := c.PostForm("token")

	err := h.service.TokenVerify(email, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Error: err.Error()})
		return
	}

	response := dto.Response{Message: "Email sucessfully verified!"}
	c.JSON(http.StatusOK, response)

}