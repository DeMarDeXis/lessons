package storage

import (
	"courses/internal/domain"
	"courses/internal/storage/postgres"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Lesson interface {
	CreateLesson(lesson *domain.Lesson) (int, error)
	GetLessonByName(name string) (*domain.Lesson, error)
	GetLessonByID(id int) (*domain.Lesson, error)
	SendLessonForMarking(lessonID int) error
	GetAllDoneLesson() (*[]domain.Lesson, error)
	UpdateLessonStatus(name int, status string) error
}

type Storage struct {
	Lesson
}

// TODO: init storage
func NewStorage(db *sqlx.DB, logg *slog.Logger) *Storage {
	return &Storage{
		Lesson: postgres.NewLessonStorage(db, logg),
	}
}
