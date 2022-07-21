package service

import (
	gym "github.com/igorgofman/gms-app"
	"github.com/igorgofman/gms-app/pkg/repository"
)

type VisitService struct {
	repo repository.Visit
}

func NewVisitService(repo repository.Visit) *VisitService {
	return &VisitService{repo: repo}
}

func (s *VisitService) Create(visitorId int) (int, error) {
	return s.repo.Create(visitorId)
}

func (s *VisitService) GetAll() ([]gym.Visit, error) {
	return s.repo.GetAll()
}

func (s *VisitService) GetById(visitId int) (gym.Visit, error) {
	return s.repo.GetById(visitId)
}

func (s *VisitService) GetAllById(visitorId int) ([]gym.Visit, error) {
	return s.repo.GetAllById(visitorId)
}

func (s *VisitService) Delete(visitId int) error {
	return s.repo.Delete(visitId)
}

func (s *VisitService) Update(visitId int, input gym.UpdateVisitInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(visitId, input)
}
