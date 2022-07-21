package repository

import (
	"errors"
	"fmt"
	"strings"

	gym "github.com/igorgofman/gms-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (r *InfoPostgres) GetAllInstructors() ([]gym.DataToPrintInstructor, error) {
	var infoInstructors []gym.DataToPrintInstructor

	query := fmt.Sprintf(`SELECT inf.id, inf.first_name, inf.last_name, inf.middle_name, inf.relationship, inf.phone, inf.date_of_birth, inf.date_of_registry, inst.hire_date, inst.salary FROM %s inf INNER JOIN %s inst on inst.info_id = inf.id WHERE inf.relationship = 'instructor'`, infoTable, instructorsTable)
	err := r.db.Select(&infoInstructors, query)

	return infoInstructors, err
}

func (r *InfoPostgres) GetAllMembers() ([]gym.DataToPrintMember, error) {
	var infoMembers []gym.DataToPrintMember

	query := fmt.Sprintf(`SELECT inf.id, inf.first_name, inf.last_name, inf.middle_name, inf.relationship, inf.phone, inf.date_of_birth, inf.date_of_registry, mem.membership_id, mem.expires_at FROM %s inf INNER JOIN %s mem on mem.info_id = inf.id WHERE inf.relationship = 'member'`, infoTable, membersTable)
	err := r.db.Select(&infoMembers, query)

	return infoMembers, err
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

func (r *InfoPostgres) Update(infoId int, input gym.UpdateInfoInput) error {
	var info gym.Info

	row := fmt.Sprintf(`SELECT relationship FROM %s WHERE id=$1`, infoTable)
	_ = r.db.Get(&info, row, infoId)

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.First_Name != nil {
		setValues = append(setValues, fmt.Sprintf("first_name=$%d", argId))
		args = append(args, *input.First_Name)
		argId++
	}

	if input.Last_Name != nil {
		setValues = append(setValues, fmt.Sprintf("last_name=$%d", argId))
		args = append(args, *input.Last_Name)
		argId++
	}

	if input.Middle_Name != nil {
		setValues = append(setValues, fmt.Sprintf("middle_name=$%d", argId))
		args = append(args, *input.Middle_Name)
		argId++
	}

	if input.Relationship != nil {
		setValues = append(setValues, fmt.Sprintf("relationship=$%d", argId))
		args = append(args, *input.Relationship)
		argId++
	}

	if input.Phone != nil {
		setValues = append(setValues, fmt.Sprintf("phone=$%d", argId))
		args = append(args, *input.Phone)
		argId++
	}

	if input.Date_of_birth != nil {
		setValues = append(setValues, fmt.Sprintf("date_of_birth=$%d", argId))
		args = append(args, *input.Date_of_birth)
		argId++
	}

	switch info.Relationship {
	case "member":
		setValuesMember := make([]string, 0)
		args := make([]interface{}, 0)
		argId := 1

		if input.MembershipId != nil {
			setValuesMember = append(setValuesMember, fmt.Sprintf("membership_id=$%d", argId))
			args = append(args, *input.MembershipId)
			argId++
		}

		if input.Expires_at != nil {
			setValuesMember = append(setValuesMember, fmt.Sprintf("expires_at=$%d", argId))
			args = append(args, *input.Expires_at)
			argId++
		}

		setQuery := strings.Join(setValuesMember, ", ")

		query := fmt.Sprintf("UPDATE %s SET %s WHERE info_id=$%d ", membersTable, setQuery, argId)

		args = append(args, infoId)

		logrus.Debugf("updateQueryMember: %s", query)
		logrus.Debugf("argsMember: %s", args)

		_, err := r.db.Exec(query, args...)

		if err != nil {
			return err
		}

	case "instructor":
		setValuesInstructor := make([]string, 0)
		args := make([]interface{}, 0)
		argId := 1

		if input.Salary != nil {
			setValuesInstructor = append(setValuesInstructor, fmt.Sprintf("salary=$%d", argId))
			args = append(args, *input.Salary)
			argId++
		}

		setQuery := strings.Join(setValuesInstructor, ", ")

		query := fmt.Sprintf("UPDATE %s SET %s WHERE info_id=$%d ", instructorsTable, setQuery, argId)

		args = append(args, infoId)

		logrus.Debugf("updateQueryInstructor: %s", query)
		logrus.Debugf("argsInstructor: %s", args)

		_, err := r.db.Exec(query, args...)

		if err != nil {
			return err
		}
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d ", infoTable, setQuery, argId)

	args = append(args, infoId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
