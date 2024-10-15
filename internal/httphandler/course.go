package httphandler

import (
	"courses/internal/domain"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (h *Handler) createCourse(w http.ResponseWriter, r *http.Request) {
	var input domain.Course

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}

	//TODO: init user from 'user service'
	id, err := h.service.Course.CreateCourse(&input, 1)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"id": id}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) getCourseByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid ID")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid ID")
		return
	}

	course, err := h.service.Course.GetCourseByID(idInt)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to get course by id: %s, %e", id, err))
		return
	}
	if course == nil {
		newErrorResponse(w, h.logg, http.StatusNotFound, "course not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(course)
}

func (h *Handler) updateCourse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid ID")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid ID")
		return
	}

	var input domain.UpdateCourse
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}
	if err := input.Validate(); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Course.UpdateCourse(idInt, &input); err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to update course: %s, %e", id, err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

type AllCoursesResponse struct {
	Courses []domain.Course `json:"courses"`
}

func (h *Handler) getAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := h.service.Course.GetAllCourses()
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to get all courses: %e", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := AllCoursesResponse{
		Courses: *courses,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) getAllCoursesByTeacher(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid ID")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid ID")
		return
	}

	courses, err := h.service.Course.GetAllCoursesByTeacher(idInt)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to get all courses by teacher: %s, %e", id, err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := AllCoursesResponse{
		Courses: *courses,
	}
	json.NewEncoder(w).Encode(response)
}
