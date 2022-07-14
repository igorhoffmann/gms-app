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

func (s *InfoService) GetAll() ([]gym.DataToPrintInfo, error) {
	return s.repo.GetAll()
}

func (s *InfoService) GetById(infoId int) (interface{}, error) {
	return s.repo.GetById(infoId)
}

func (s *InfoService) Delete(infoId int) error {
	return s.repo.Delete(infoId)
}

// func (s *InfoService) Update(infoId, listId int, input gym.UpdateListInput) error {
// 	if err := input.Validate(); err != nil {
// 		return err
// 	}
// 	return s.repo.Update(infoId, listId, input)
// }
