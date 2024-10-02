package service

import (
	"courses/internal/domain"
	"courses/internal/storage"
)

type Lesson interface {
	CreateLesson(lesson *domain.Lesson) (int, error)
	GetLessonByName(name string) (*domain.Lesson, error)
	GetLessonByID(id int) (*domain.Lesson, error)
	SendLessonForMarking(lessonID int) error
	GetAllDoneLesson() (*[]domain.Lesson, error)
	GetAllDoneLessonByCourse(course int) (*[]domain.Lesson, error)
	UpdateLessonStatus(name int, status string) error
}

type Service struct {
	Lesson
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Lesson: NewLessonService(storage.Lesson),
	}
}
