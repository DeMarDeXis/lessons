package postgres

import (
	"courses/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Lesson struct {
	db     *sqlx.DB
	logger *slog.Logger
}

func NewLessonStorage(db *sqlx.DB, logg *slog.Logger) *Lesson {
	return &Lesson{
		db:     db,
		logger: logg,
	}
}

func (t *Lesson) CreateLesson(lesson *domain.Lesson) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		t.logger.Error("failed to begin transaction", slog.String("err", err.Error()))
		return 0, err
	}

	var lessonID int
	q := fmt.Sprintf(`INSERT INTO %s (course_id, lesson_name, lesson_description) VALUES ($1, $2, $3) RETURNING lesson_id`, lessonsTable)
	row := tx.QueryRow(q, lesson.CourseID, lesson.Name, lesson.Description)
	if err := row.Scan(&lessonID); err != nil {
		tx.Rollback()
		t.logger.Error("failed to scan lesson id", slog.String("err", err.Error()))
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		t.logger.Error("failed to commit transaction", slog.String("err", err.Error()))
		return 0, err
	}

	return lessonID, nil
}

func (t *Lesson) GetLessonByName(name string) (*domain.Lesson, error) {
	var lesson domain.Lesson
	//TO DO: find out how to get lesson by *
	q := fmt.Sprintf(`SELECT * FROM %s WHERE lesson_name = $1`, lessonsTable)
	if err := t.db.Get(&lesson, q, name); err != nil {
		t.logger.Error("failed to get lesson by name", slog.String("err", err.Error()))
		return nil, err
	}

	if lesson.Status == nil {
		lesson.Status = new(string)
		*lesson.Status = "not done"
	}

	return &lesson, nil
}

func (t *Lesson) GetLessonByID(id int) (*domain.Lesson, error) {
	var lesson domain.Lesson
	q := fmt.Sprintf(`SELECT lesson_name, course_id, lesson_description  FROM %s WHERE lesson_id = $1`, lessonsTable)
	if err := t.db.Get(&lesson, q, id); err != nil {
		t.logger.Error("failed to get lesson by id", slog.String("err", err.Error()))
		return nil, err
	}
	return &lesson, nil
}

func (t *Lesson) GetAllDoneLesson() (*[]domain.Lesson, error) {
	var lessons []domain.Lesson
	q := fmt.Sprintf(`SELECT lesson_name, course_id, lesson_description  FROM %s WHERE status = $1`, lessonsTable)
	err := t.db.Select(&lessons, q, "done")
	if err != nil {
		t.logger.Error("failed to get all done lessons", slog.String("err", err.Error()))
		return nil, err
	}

	return &lessons, nil
}

func (t *Lesson) SendLessonForMarking(lessonID int) error {
	q := fmt.Sprintf(`INSERT INTO %s (teacher_id, lesson_id, status, created_at, updated_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, teachersChecklistTable)
	_, err := t.db.Exec(q, 1, lessonID, "done")
	if err != nil {
		t.logger.Error("failed to send lesson for marking", slog.String("err", err.Error()))
		return err
	}
	return nil
}

func (t *Lesson) UpdateLessonStatus(name int, status string) error {
	q := fmt.Sprintf(`UPDATE %s SET status = $1 WHERE lesson_id = $2`, lessonsTable)
	_, err := t.db.Exec(q, status, name)
	if err != nil {
		t.logger.Error("failed to update lesson status", slog.String("err", err.Error()))
		return err
	}
	return nil
}
