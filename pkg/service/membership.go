package service

import (
	gym "github.com/igorgofman/GMS-app"
	"github.com/igorgofman/GMS-app/pkg/repository"
)

type MembershipService struct {
	repo repository.Membership
}

func NewMembershipService(repo repository.Membership) *MembershipService {
	return &MembershipService{repo: repo}
}

func (s *MembershipService) Create(membership gym.Membership) (int, error) {
	return s.repo.Create(membership)
}

func (s *MembershipService) GetAll() ([]gym.Membership, error) {
	return s.repo.GetAll()
}

func (s *MembershipService) GetById(membershipId int) (gym.Membership, error) {
	return s.repo.GetById(membershipId)
}

func (s *MembershipService) Delete(membershipId int) error {
	return s.repo.Delete(membershipId)
}

func (s *MembershipService) Update(membershipId int, input gym.UpdateMembershipInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(membershipId, input)
}
