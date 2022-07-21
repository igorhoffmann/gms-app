package repository

import (
	"fmt"
	"strings"

	gym "github.com/igorgofman/gms-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type MembershipPostgres struct {
	db *sqlx.DB
}

func NewMembershipPostgres(db *sqlx.DB) *MembershipPostgres {
	return &MembershipPostgres{db: db}
}

func (r *MembershipPostgres) Create(membership gym.Membership) (int, error) {
	var id int
	createMembershipQuery := fmt.Sprintf("INSERT INTO %s (title, price, duration, instructor_id) VALUES ($1, $2, $3, $4) RETURNING id", membershipsTable)
	row := r.db.QueryRow(createMembershipQuery, membership.Title, membership.Price, membership.Duration, membership.InstructorId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *MembershipPostgres) GetAll() ([]gym.Membership, error) {
	var memberships []gym.Membership

	query := fmt.Sprintf("SELECT * FROM %s", membershipsTable)
	err := r.db.Select(&memberships, query)

	return memberships, err
}

func (r *MembershipPostgres) GetById(membershipId int) (gym.Membership, error) {
	var membership gym.Membership

	row := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, membershipsTable)
	err := r.db.Get(&membership, row, membershipId)

	return membership, err
}

func (r *MembershipPostgres) Delete(membershipId int) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", membershipsTable)
	_, err := r.db.Exec(query, membershipId)

	return err
}

func (r *MembershipPostgres) Update(membershipId int, input gym.UpdateMembershipInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	if input.Duration != nil {
		setValues = append(setValues, fmt.Sprintf("duration=$%d", argId))
		args = append(args, *input.Duration)
		argId++
	}

	if input.InstructorId != nil {
		setValues = append(setValues, fmt.Sprintf("instructor_id=$%d", argId))
		args = append(args, *input.InstructorId)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", membershipsTable, setQuery, argId)

	args = append(args, membershipId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
