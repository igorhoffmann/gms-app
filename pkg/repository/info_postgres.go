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

	// fmt.Printf("2: %v\n", instructor.Salary)

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

// func (r *InfoPostgres) GetAll(infoId int) ([]gym.Info, error) {
// 	var lists []gym.Info

// 	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", todoListsTable, usersListsTable)
// 	err := r.db.Select(&lists, query, infoId)

// 	return lists, err
// }

// func (r *InfoPostgres) GetById(infoId, listId int) (gym.Info, error) {
// 	var info gym.Info

// 	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
// INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`, todoListsTable, usersListsTable)
// 	err := r.db.Get(&info, query, infoId, listId)

// 	return info, err
// }

// func (r *InfoPostgres) Delete(infoId, listId int) error {
// 	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2", todoListsTable, usersListsTable)
// 	_, err := r.db.Exec(query, infoId, listId)

// 	return err
// }

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
