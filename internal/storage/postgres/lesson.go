package postgres

import (
	"courses/internal/domain"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"strings"
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
	q := fmt.Sprintf(`INSERT INTO %s (lesson_name, lesson_description, lesson_type) VALUES ($1, $2, $3) RETURNING lesson_id`, lessonsTable)
	row := tx.QueryRow(q, lesson.Name, lesson.Description, lesson.LessonType)
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
	q := fmt.Sprintf(`SELECT * FROM %s WHERE lesson_name = $1`, lessonsTable)
	err := t.db.Get(&lesson, q, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			t.logger.Info("no lesson found with the given name", slog.String("name", name))
			return nil, fmt.Errorf("lesson not found: %s", name)
		}
		t.logger.Error("failed to get lesson by name", slog.String("err", err.Error()))
		return nil, err
	}

	return &lesson, nil
}

func (t *Lesson) GetLessonByID(id int) (*domain.Lesson, error) {
	var lesson domain.Lesson
	q := fmt.Sprintf(`SELECT * FROM %s WHERE lesson_id = $1`, lessonsTable)
	if err := t.db.Get(&lesson, q, id); err != nil {
		t.logger.Error("failed to get lesson by id", slog.String("err", err.Error()))
		return nil, err
	}
	return &lesson, nil
}

func (t *Lesson) UploadFile(lessonID int, fileName string, fileData []byte) error {
	q := fmt.Sprintf(`UPDATE %s SET lesson_file_name = $1, lesson_file_content = $2 WHERE lesson_id = $3`, lessonsTable)
	_, err := t.db.Exec(q, fileName, fileData, lessonID)
	if err != nil {
		t.logger.Error("failed to upload file", slog.String("err", err.Error()))
		return err
	}
	return nil
}

func (t *Lesson) SendLessonForMarking(lessonID int) error {
	tx, err := t.db.Begin()
	if err != nil {
		t.logger.Error("failed to begin transaction", slog.String("err", err.Error()))
		return err
	}

	q := fmt.Sprintf(`SELECT lesson_file_content FROM %s WHERE lesson_id = $1`, lessonsTable)
	var fileData []byte
	err = tx.QueryRow(q, lessonID).Scan(&fileData)
	if err != nil {
		tx.Rollback()
		t.logger.Error("failed to get file data", slog.String("err", err.Error()))
		return err
	}

	q = fmt.Sprintf(`UPDATE %s SET lesson_status = $1 WHERE lesson_id = $2`, lessonsTable)
	_, err = tx.Exec(q, "send", lessonID)
	if err != nil {
		tx.Rollback()
		t.logger.Error("failed to update lesson status", slog.String("err", err.Error()))
		return err
	}

	q = fmt.Sprintf(`INSERT INTO %s (teacher_id, lesson_id, homework, created_at, updated_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, teachersChecklistTable)
	//TODO: fix this userID
	result, err := t.db.Exec(q, 1, lessonID, fileData)
	if err != nil {
		t.logger.Error("failed to send lesson for marking", slog.String("err", err.Error()))
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		t.logger.Error("failed to get rows affected", slog.String("err", err.Error()))
		return err
	}

	if rowsAffected == 0 {
		tx.Rollback()
		t.logger.Warn("no rows were updated", slog.Int("lessonID", lessonID))
		return fmt.Errorf("no rows were updated for lessonID %d", lessonID)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		t.logger.Error("failed to commit transaction", slog.String("err", err.Error()))
		return err
	}
	return nil
}

func (t *Lesson) GetAllDoneLesson() (*[]domain.Lesson, error) {
	var lessons []domain.Lesson
	q := fmt.Sprintf(`SELECT lesson_name, lesson_description  FROM %s WHERE lesson_status = $1`, lessonsTable)
	err := t.db.Select(&lessons, q, "send")
	if err != nil {
		t.logger.Error("failed to get all send lessons", slog.String("err", err.Error()))
		return nil, err
	}

	return &lessons, nil
}

// TODO: fix it
func (t *Lesson) GetAllDoneLessonByCourse(course int) (*[]domain.Lesson, error) {
	var lessons []domain.Lesson
	q := fmt.Sprintf(`SELECT lesson_name, lesson_description  FROM %s WHERE course_id = $1 AND status = $2`, lessonsTable)
	err := t.db.Select(&lessons, q, course, "done")
	if err != nil {
		t.logger.Error("failed to get all done lessons by course", slog.String("err", err.Error()))
		return nil, err
	}

	return &lessons, nil
}

func (t *Lesson) UpdateLesson(id int, lessonData *domain.UpdateLesson) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if lessonData.LessonType != nil {
		setValues = append(setValues, fmt.Sprintf("lesson_type=$%d", argID))
		args = append(args, *lessonData.LessonType)
		argID++
	}

	if lessonData.Name != nil {
		setValues = append(setValues, fmt.Sprintf("lesson_name=$%d", argID))
		args = append(args, *lessonData.Name)
		argID++
	}

	if lessonData.Description != nil {
		setValues = append(setValues, fmt.Sprintf("lesson_description=$%d", argID))
		args = append(args, *lessonData.Description)
		argID++
	}

	setQ := strings.Join(setValues, ", ")

	q := fmt.Sprintf(`UPDATE %s SET %s WHERE lesson_id = $%d`, lessonsTable, setQ, argID)
	args = append(args, id)
	_, err := t.db.Exec(q, args...)
	if err != nil {
		t.logger.Error("failed to update lesson", slog.String("err", err.Error()))
		return err
	}
	return nil
}
