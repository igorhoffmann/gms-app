package repository

import (
	"fmt"
	"strings"

	gym "github.com/igorgofman/GMS-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type VisitPostgres struct {
	db *sqlx.DB
}

func NewVisitPostgres(db *sqlx.DB) *VisitPostgres {
	return &VisitPostgres{db: db}
}

func (r *VisitPostgres) Create(visitorId int) (int, error) {
	var id int
	createMembershipQuery := fmt.Sprintf("INSERT INTO %s (visitor_id) VALUES ($1) RETURNING id", visitsTable)
	row := r.db.QueryRow(createMembershipQuery, visitorId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *VisitPostgres) GetAll() ([]gym.Visit, error) {
	var visits []gym.Visit

	query := fmt.Sprintf("SELECT * FROM %s", visitsTable)
	err := r.db.Select(&visits, query)

	return visits, err
}

func (r *VisitPostgres) GetById(visitId int) (gym.Visit, error) {
	var visit gym.Visit

	row := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, visitsTable)
	err := r.db.Get(&visit, row, visitId)

	return visit, err
}

func (r *VisitPostgres) GetAllById(visitorId int) ([]gym.Visit, error) {
	var visits []gym.Visit

	query := fmt.Sprintf(`SELECT * FROM %s WHERE visitor_id=$1`, visitsTable)
	err := r.db.Select(&visits, query, visitorId)

	return visits, err
}

func (r *VisitPostgres) Delete(visitId int) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", visitsTable)
	_, err := r.db.Exec(query, visitId)

	return err
}

func (r *VisitPostgres) Update(visitId int, input gym.UpdateVisitInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.InfoId != nil {
		setValues = append(setValues, fmt.Sprintf("visitor_id=$%d", argId))
		args = append(args, *input.InfoId)
		argId++
	}

	if input.Came_at != nil {
		setValues = append(setValues, fmt.Sprintf("came_at=$%d", argId))
		args = append(args, *input.Came_at)
		argId++
	}

	if input.Left_at != nil {
		setValues = append(setValues, fmt.Sprintf("left_at=$%d", argId))
		args = append(args, *input.Left_at)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", visitsTable, setQuery, argId)

	args = append(args, visitId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
