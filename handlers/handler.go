package handlers

import "email-verification/services"

type compHanders struct {
	service services.CompServices
}

func NewCompHandlers(s services.CompServices) *compHanders {
	return &compHanders{
		service: s,
	}
}