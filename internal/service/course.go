package service

import (
	"courses/internal/domain"
	"courses/internal/storage"
)

type CourseService struct {
	storage storage.Courses
}

func NewCourseService(storage storage.Courses) *CourseService {
	return &CourseService{storage: storage}
}

func (s *CourseService) CreateCourse(course *domain.Course, userID int) (int, error) {
	return s.storage.CreateCourse(course, userID)
}

func (s *CourseService) GetCourseByID(id int) (*domain.Course, error) {
	return s.storage.GetCourseByID(id)
}

func (s *CourseService) UpdateCourse(id int, course *domain.UpdateCourse) error {
	return s.storage.UpdateCourse(id, course)
}

func (s *CourseService) GetAllCourses() (*[]domain.Course, error) {
	return s.storage.GetAllCourses()
}

func (s *CourseService) GetAllCoursesByTeacher(userID int) (*[]domain.Course, error) {
	return s.storage.GetAllCoursesByTeacher(userID)
}
