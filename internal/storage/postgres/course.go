package postgres

import (
	"courses/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"strings"
)

type Course struct {
	db     *sqlx.DB
	logger *slog.Logger
}

func NewCourseStorage(db *sqlx.DB, logg *slog.Logger) *Course {
	return &Course{
		db:     db,
		logger: logg,
	}
}

func (t *Course) CreateCourse(course *domain.Course, userID int) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		t.logger.Error("failed to begin transaction", slog.String("err", err.Error()))
		return 0, err
	}

	var courseID int
	q := fmt.Sprintf(`INSERT INTO %s (name, description, owner_id) VALUES ($1, $2, $3) RETURNING id`, coursesTable)

	row := tx.QueryRow(q, course.Name, course.Description, userID)
	if err := row.Scan(&courseID); err != nil {
		tx.Rollback()
		t.logger.Error("failed to scan course id", slog.String("err", err.Error()))
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		t.logger.Error("failed to commit transaction", slog.String("err", err.Error()))
		return 0, err
	}

	return courseID, nil
}

func (t *Course) GetCourseByID(id int) (*domain.Course, error) {
	var course domain.Course

	q := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, coursesTable)
	if err := t.db.Get(&course, q, id); err != nil {
		t.logger.Error("failed to get course by id", slog.String("err", err.Error()))
		return nil, err
	}
	return &course, nil
}

func (t *Course) UpdateCourse(id int, course *domain.UpdateCourse) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if course.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argID))
		args = append(args, *course.Name)
		argID++
	}

	if course.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *course.Description)
		argID++
	}

	setValues = append(setValues, "updated_at=NOW()")

	setQuery := strings.Join(setValues, ", ")

	q := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d`, coursesTable, setQuery, argID)
	args = append(args, id)
	_, err := t.db.Exec(q, args...)
	if err != nil {
		t.logger.Error("failed to update course", slog.String("err", err.Error()))
		return err
	}
	return nil
}

func (t *Course) GetAllCourses() (*[]domain.Course, error) {
	var courses []domain.Course
	q := fmt.Sprintf(`SELECT * FROM %s`, coursesTable)
	if err := t.db.Select(&courses, q); err != nil {
		t.logger.Error("failed to get all courses", slog.String("err", err.Error()))
		return nil, err
	}
	return &courses, nil
}

func (t *Course) GetAllCoursesByTeacher(userID int) (*[]domain.Course, error) {
	var courses []domain.Course
	q := fmt.Sprintf(`SELECT * FROM %s WHERE owner_id = $1`, coursesTable)
	if err := t.db.Select(&courses, q, userID); err != nil {
		t.logger.Error("failed to get all courses by teacher", slog.String("err", err.Error()))
		return nil, err
	}
	return &courses, nil
}
