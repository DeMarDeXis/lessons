package service

import (
	"courses/internal/domain"
	"courses/internal/storage"
	"log"
	"os"
	"path/filepath"
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

func (s *LessonService) UpdateLesson(id int, lessonData *domain.UpdateLesson) error {
	return s.storage.UpdateLesson(id, lessonData)
}

func (s *LessonService) UploadFile(lessonID int, fileName string, fileData []byte) error {
	//dir := "./lessonFiles"
	//if err := os.MkdirAll(dir, os.ModePerm); err != nil {
	//	return err
	//}

	//filePath := filepath.Join(dir, fileName)
	filePath := filepath.Join("lessonFiles", fileName)

	defer func() {
		err := os.Remove(filePath)
		if err != nil {
			log.Println("Ошибка при удалении файла:", err)
		}
	}()

	//if err := os.WriteFile(filePath, fileData, 0644); err != nil {
	//	return err
	//}

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = s.storage.UploadFile(lessonID, fileName, fileData)
	if err != nil {
		return err
	}

	return nil
}

func (s *LessonService) SendLessonForMarking(lessonID int) error {
	return s.storage.SendLessonForMarking(lessonID)
}

func (s *LessonService) GetAllDoneLesson() (*[]domain.Lesson, error) {
	return s.storage.GetAllDoneLesson()
}

func (s *LessonService) GetAllDoneLessonByCourse(course int) (*[]domain.Lesson, error) {
	return s.storage.GetAllDoneLessonByCourse(course)
}
