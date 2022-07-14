package repository

import (
	"errors"
	"fmt"

	// "strings"

	gym "github.com/igorgofman/GMS-app"
	"github.com/jmoiron/sqlx"
	// "github.com/sirupsen/logrus"
)

type InfoPostgres struct {
	db *sqlx.DB
}

func NewInfoPostgres(db *sqlx.DB) *InfoPostgres {
	return &InfoPostgres{db: db}
}

func (r *InfoPostgres) Create(info gym.Info, member gym.Member, instructor gym.Instructor) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createInfoQuery := fmt.Sprintf("INSERT INTO %s (first_name, last_name, middle_name, relationship, phone, date_of_birth) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", infoTable)
	row := tx.QueryRow(createInfoQuery, info.First_Name, info.Last_Name, info.Middle_Name, info.Relationship, info.Phone, info.Date_of_birth)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	switch info.Relationship {
	case "member":
		createMemberQuery := fmt.Sprintf("INSERT INTO %s (info_id, membership_id, expires_at) VALUES ($1, $2, $3)", membersTable)
		_, err = tx.Exec(createMemberQuery, id, info.Member.MembershipId, info.Member.Expires_at)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	case "instructor":
		createInstructorQuery := fmt.Sprintf("INSERT INTO %s (info_id, salary) VALUES ($1, $2)", instructorsTable)
		_, err = tx.Exec(createInstructorQuery, id, info.Instructor.Salary)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	default:
		tx.Rollback()
		return 0, errors.New("pq:info: relationship input error")
	}

	return id, tx.Commit()
}

func (r *InfoPostgres) GetAll() ([]gym.DataToPrintInfo, error) {
	var infos []gym.DataToPrintInfo

	query := fmt.Sprintf("SELECT id, first_name, last_name, relationship, phone, date_of_registry FROM %s", infoTable)
	err := r.db.Select(&infos, query)

	return infos, err
}

func (r *InfoPostgres) GetById(infoId int) (interface{}, error) {
	var info gym.Info

	row := fmt.Sprintf(`SELECT relationship FROM %s WHERE id=$1`, infoTable)
	_ = r.db.Get(&info, row, infoId)

	switch info.Relationship {
	case "member":
		var infoMember gym.DataToPrintMember
		query := fmt.Sprintf(`SELECT inf.id, inf.first_name, inf.last_name, inf.middle_name, inf.relationship, inf.phone, inf.date_of_birth, inf.date_of_registry, mem.membership_id, mem.expires_at FROM %s inf INNER JOIN %s mem on mem.info_id = inf.id WHERE inf.id = $1 AND  mem.info_id = $1`, infoTable, membersTable)
		err := r.db.Get(&infoMember, query, infoId)
		return infoMember, err
	case "instructor":
		var infoInstructor gym.DataToPrintInstructor
		query := fmt.Sprintf(`SELECT inf.id, inf.first_name, inf.last_name, inf.middle_name, inf.relationship, inf.phone, inf.date_of_birth, inf.date_of_registry, inst.hire_date, inst.salary FROM %s inf INNER JOIN %s inst on inst.info_id = inf.id WHERE inf.id = $1 AND  inst.info_id = $1`, infoTable, instructorsTable)
		err := r.db.Get(&infoInstructor, query, infoId)
		return infoInstructor, err
	default:
		return info, errors.New("pq:info: relationship input error")
	}

}

func (r *InfoPostgres) Delete(infoId int) error {
	var info gym.Info

	row := fmt.Sprintf(`SELECT relationship FROM %s WHERE id=$1`, infoTable)
	_ = r.db.Get(&info, row, infoId)

	switch info.Relationship {
	case "member":
		memrow := fmt.Sprintf("DELETE FROM %s WHERE info_id = $1", membersTable)
		_, err := r.db.Exec(memrow, infoId)
		if err != nil {
			return err
		}
	case "instructor":
		instrow := fmt.Sprintf("DELETE FROM %s WHERE info_id = $1", instructorsTable)
		_, err := r.db.Exec(instrow, infoId)
		if err != nil {
			return err
		}
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", infoTable)
	_, err := r.db.Exec(query, infoId)

	return err
}

// func (r *InfoPostgres) Update(infoId, listId int, input gym.UpdateListInput) error {
// 	setValues := make([]string, 0)
// 	args := make([]interface{}, 0)
// 	argId := 1

// 	if input.Title != nil {
// 		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
// 		args = append(args, *input.Title)
// 		argId++
// 	}

// 	if input.Description != nil {
// 		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
// 		args = append(args, *input.Description)
// 		argId++
// 	}

// 	setQuery := strings.Join(setValues, ", ")

// 	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s vl WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d", todoListsTable, setQuery, usersListsTable, argId, argId+1)

// 	args = append(args, listId, infoId)

// 	logrus.Debug("updateQuery: %s", query)
// 	logrus.Debug("args: %s", args)

// 	_, err := r.db.Exec(query, args...)
// 	return err
// }
