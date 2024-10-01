package service

import (
	"courses/internal/domain"
	"courses/internal/storage"
)

type LessonService struct {
	storage storage.Lesson
}

func NewLessonService(storage storage.Lesson) *LessonService {
	return &LessonService{storage: storage}
}

func (s *LessonService) CreateLesson(lesson *domain.Lesson) (int, error) {
	return s.storage.CreateLesson(lesson)
}

func (s *LessonService) GetLessonByName(name string) (*domain.Lesson, error) {
	return s.storage.GetLessonByName(name)
}

func (s *LessonService) GetLessonByID(id int) (*domain.Lesson, error) {
	return s.storage.GetLessonByID(id)
}

func (s *LessonService) GetAllDoneLesson() (*[]domain.Lesson, error) {
	return s.storage.GetAllDoneLesson()
}

func (s *LessonService) SendLessonForMarking(lessonID int) error {
	return s.storage.SendLessonForMarking(lessonID)
}

func (s *LessonService) UpdateLessonStatus(name int, status string) error {
	return s.storage.UpdateLessonStatus(name, status)
}
