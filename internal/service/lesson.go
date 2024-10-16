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

func (s *LessonService) CreateLesson(courseID int, lesson *domain.Lesson) (int, error) {
	return s.storage.CreateLesson(courseID, lesson)
}

func (s *LessonService) GetLessonByName(courseID int, name string) (*domain.Lesson, error) {
	return s.storage.GetLessonByName(courseID, name)
}

func (s *LessonService) GetLessonByID(courseID int, id int) (*domain.Lesson, error) {
	return s.storage.GetLessonByID(courseID, id)
}

func (s *LessonService) GetAllLessons(courseID int) (*[]domain.Lesson, error) {
	return s.storage.GetAllLessons(courseID)

}

func (s *LessonService) UpdateLesson(courseID int, id int, lessonData *domain.UpdateLesson) error {
	return s.storage.UpdateLesson(courseID, id, lessonData)
}

func (s *LessonService) UploadFile(courseID int, lessonID int, fileName string, fileData []byte) error {
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

	err = s.storage.UploadFile(courseID, lessonID, fileName, fileData)
	if err != nil {
		return err
	}

	return nil
}

func (s *LessonService) SendLessonForMarking(courseID int, lessonID int) error {
	return s.storage.SendLessonForMarking(courseID, lessonID)
}

//func (s *LessonService) GetAllDoneLesson() (*[]domain.Lesson, error) {
//	return s.storage.GetAllDoneLesson()
//}

//func (s *LessonService) GetAllDoneLessonByCourse(course int) (*[]domain.Lesson, error) {
//	return s.storage.GetAllDoneLessonByCourse(course)
//}
