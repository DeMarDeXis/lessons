package httphandler

import (
	"courses/internal/domain"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (h *Handler) createLesson(w http.ResponseWriter, r *http.Request) {
	courseIDInt, err := h.getCourseIDFromRequest(w, r)
	if err != nil {
		return
	}

	var input domain.Lesson

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Lesson.CreateLesson(courseIDInt, &input)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"id": id}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) getLessonByName(w http.ResponseWriter, r *http.Request) {
	courseIDInt, err := h.getCourseIDFromRequest(w, r)
	if err != nil {
		return
	}

	name := chi.URLParam(r, "name")
	if name == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid name")
		return
	}

	lesson, err := h.service.Lesson.GetLessonByName(courseIDInt, name)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to get lesson by name: %s, %e", name, err))
		return
	}
	if lesson == nil {
		newErrorResponse(w, h.logg, http.StatusNotFound, "lesson not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//response := map[string]interface{}{"lesson": lesson}
	json.NewEncoder(w).Encode(lesson)
}

func (h *Handler) getLessonByID(w http.ResponseWriter, r *http.Request) {
	courseIDInt, err := h.getCourseIDFromRequest(w, r)
	if err != nil {
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid id")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid id")
		return
	}
	lesson, err := h.service.Lesson.GetLessonByID(courseIDInt, idInt)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to get lesson by id: %s, %e", id, err))
		return
	}
	if lesson == nil {
		newErrorResponse(w, h.logg, http.StatusNotFound, "lesson not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//response := map[string]interface{}{"lesson": lesson}
	json.NewEncoder(w).Encode(lesson)
}

func (h *Handler) getAllLessons(w http.ResponseWriter, r *http.Request) {
	courseIDInt, err := h.getCourseIDFromRequest(w, r)
	if err != nil {
		return
	}

	lessons, err := h.service.Lesson.GetAllLessons(courseIDInt)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to get all lessons: %e", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//response := map[string]interface{}{"lessons": lessons}
	json.NewEncoder(w).Encode(lessons)
}

func (h *Handler) updateLesson(w http.ResponseWriter, r *http.Request) {
	courseIDInt, err := h.getCourseIDFromRequest(w, r)
	if err != nil {
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid id")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid id")
		return
	}

	var input domain.UpdateLesson
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}
	if err := input.Validate(); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Lesson.UpdateLesson(courseIDInt, idInt, &input); err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to update lesson: %s, %e", id, err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) uploadFile(w http.ResponseWriter, r *http.Request) {
	courseIDInt, err := h.getCourseIDFromRequest(w, r)
	if err != nil {
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid id")
		return
	}

	filename := chi.URLParam(r, "filename") + ".txt"
	if filename == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid filename")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid id")
		return
	}

	filedata := r.FormValue("content")

	if err := h.service.Lesson.UploadFile(courseIDInt, idInt, filename, []byte(filedata)); err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to upload file: %s, %e", id, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully"))
}

func (h *Handler) sendLessonForMarking(w http.ResponseWriter, r *http.Request) {
	courseIDInt, err := h.getCourseIDFromRequest(w, r)
	if err != nil {
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid id")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.service.Lesson.SendLessonForMarking(courseIDInt, idInt); err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to send lesson for marking: %s, %e", id, err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getCourseIDFromRequest(w http.ResponseWriter, r *http.Request) (int, error) {
	courseID := chi.URLParam(r, "course_id")
	if courseID == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid course ID")
		return 0, fmt.Errorf("empty course ID")
	}

	courseIDInt, err := strconv.Atoi(courseID)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid course ID")
		return 0, fmt.Errorf("invalid course ID: %w", err)
	}

	return courseIDInt, nil
}

//TODO: do it
//func (h *Handler) getAllDoneLessons(w http.ResponseWriter, r *http.Request) {
//	lists, err := h.service.Lesson.GetAllDoneLesson()
//	if err != nil {
//		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to get all done lessons: %e", err))
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	//response := GetAllDoneLesson{
//	//	Data: lists,
//	//}
//	json.NewEncoder(w).Encode(lists)
//}

//TODO: do it
//func (h *Handler) getAllDoneLessonsByCourse(w http.ResponseWriter, r *http.Request) {
//	course := chi.URLParam(r, "course")
//	if course == "" {
//		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid course")
//		return
//	}
//
//	courseInt, err := strconv.Atoi(course)
//	if err != nil {
//		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid course")
//		return
//	}
//
//	lists, err := h.service.Lesson.GetAllDoneLessonByCourse(courseInt)
//	if err != nil {
//		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to get all done lessons: %e", err))
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(lists)
//}
