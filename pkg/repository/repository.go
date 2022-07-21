package repository

import (
	"github.com/igorgofman/gms-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user gym.SysUser) (int, error)
	GetUser(username, password string) (gym.SysUser, error)
}

type Info interface {
	Create(info gym.Info, member gym.Member, instructor gym.Instructor) (int, error)
	GetAll() ([]gym.DataToPrintInfo, error)
	GetAllInstructors() ([]gym.DataToPrintInstructor, error)
	GetAllMembers() ([]gym.DataToPrintMember, error)
	GetById(infoId int) (interface{}, error)
	Delete(infoId int) error
	Update(infoId int, input gym.UpdateInfoInput) error
}

type Membership interface {
	Create(membership gym.Membership) (int, error)
	GetAll() ([]gym.Membership, error)
	GetById(membershipId int) (gym.Membership, error)
	Delete(membershipId int) error
	Update(membershipId int, input gym.UpdateMembershipInput) error
}

type Visit interface {
	Create(visitorId int) (int, error)
	GetAll() ([]gym.Visit, error)
	GetById(visitId int) (gym.Visit, error)
	GetAllById(visitorId int) ([]gym.Visit, error)
	Delete(visitId int) error
	Update(visitId int, input gym.UpdateVisitInput) error
}

type Repository struct {
	Authorization
	Info
	Membership
	Visit
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Info:          NewInfoPostgres(db),
		Membership:    NewMembershipPostgres(db),
		Visit:         NewVisitPostgres(db),
	}
}
