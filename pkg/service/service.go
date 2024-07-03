package service

import "main-server/pkg/repository"

type Search struct {
}

type Service struct {
	Search
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
