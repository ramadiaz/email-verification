package services

import "email-verification/repositories"

type CompServices interface {
}

type compServices struct {
	repo repositories.CompRepositories
}

func NewServices(r repositories.CompRepositories) *compServices {
	return &compServices{
		repo: r,
	}
}