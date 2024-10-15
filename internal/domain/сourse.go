package domain

type Course struct {
	CourseID    int    `json:"course_id" db:"course_id"`
	Name        string `json:"name" db:"course_name"`
	Description string `json:"description" db:"course_description"`
	//Status string `json:"status" db:"status"`

}
