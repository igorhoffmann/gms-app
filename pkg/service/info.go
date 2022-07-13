package service

import (
	gym "github.com/igorgofman/GMS-app"
	"github.com/igorgofman/GMS-app/pkg/repository"
)

type InfoService struct {
	repo repository.Info
}

func NewInfoService(repo repository.Info) *InfoService {
	return &InfoService{repo: repo}
}

func (s *InfoService) Create(info gym.Info, member gym.Member, instructor gym.Instructor) (int, error) {
	return s.repo.Create(info, member, instructor)
}

// func (s *InfoService) GetAll(infoId int) ([]gym.Info, error) {
// 	return s.repo.GetAll(infoId)
// }

// func (s *InfoService) GetById(infoId, listId int) (gym.Info, error) {
// 	return s.repo.GetById(infoId, listId)
// }

// func (s *InfoService) Delete(infoId, listId int) error {
// 	return s.repo.Delete(infoId, listId)
// }

// func (s *InfoService) Update(infoId, listId int, input gym.UpdateListInput) error {
// 	if err := input.Validate(); err != nil {
// 		return err
// 	}
// 	return s.repo.Update(infoId, listId, input)
// }
