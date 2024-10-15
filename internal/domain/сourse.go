package domain

import (
	"errors"
	"fmt"
)

type Course struct {
	CourseID    int    `json:"course_id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"desc" db:"description"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
	OwnerID     int    `json:"owner_id" db:"owner_id"`
}

type UpdateCourse struct {
	Name        *string `json:"name" db:"name"`
	Description *string `json:"desc" db:"description"`
}

func (c UpdateCourse) Validate() error {
	if c.Name == nil && c.Description == nil {
		return errors.New(fmt.Sprintf("invalid course id: %d or description: %s", c.Name, c.Description))
	}
	return nil
}
