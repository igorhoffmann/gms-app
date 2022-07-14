package service

import (
	"github.com/igorgofman/GMS-app"
	"github.com/igorgofman/GMS-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user gym.SysUser) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Info interface {
	Create(info gym.Info, member gym.Member, instructor gym.Instructor) (int, error)
	GetAll() ([]gym.DataToPrintInfo, error)
	GetById(infoId int) (interface{}, error)
	// Delete(userId, infoId int) error
	// Update(userId, infoId int, input gym.UpdateInfoInput) error
}

type Membership interface {
	// Create(userId, membership gym.Membership) (int, error)
	// GetAll(userId int) ([]gym.Membership, error)
	// GetById(userId, membershipId int) (gym.Membership, error)
	// Delete(userId, membershipId int) error
	// Update(userId, membershipId int, input gym.UpdateMembershipInput) error
}

type Visit interface {
	// Create(userId, infoId int, visit gym.Visit) (int, error)
	// GetAll(userId, infoId int) ([]gym.Visit, error)
	// GetById(userId, visitId int) (gym.Visit, error)
	// Delete(userId, visitId int) error
	// Update(userId, visitId int, input gym.UpdateVisitInput) error
}

type Service struct {
	Authorization
	Info
	Membership
	Visit
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Info:          NewInfoService(repos.Info),
		// Membership:      NewMembershipService(repos.Membership),
		// Visit:      NewVisitService(repos.Visit, repos.Info),
	}
}
