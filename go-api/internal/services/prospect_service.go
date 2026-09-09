package services

import (
	"go-api/internal/models"
	"go-api/internal/repositories"
)

type ProspectService struct {
	Repo *repositories.ProspectRepository
}

func (s *ProspectService) GetRowCount() (models.ProspectCount, error) {
	return s.Repo.FetchAll()
}
