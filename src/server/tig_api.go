package server

import (
	"time"
)

type (
	Assignment struct {
		ID             int        `json:"id"`
		CreatorID      int        `json:"creator_id"`
		AssignedID     int        `json:"assigned_id"`
		StatusID       int        `json:"status_id"`
		TaskID         int        `json:"task_id"`
		CreatedAt      *time.Time `json:"created_at"`
		LastEdit       *time.Time `json:"last_edit"`
		DueDate        *time.Time `json:"due_date"`
		CompletionDate *time.Time `json:"completion_date"`
		Title          string     `json:"title"`
		Description    string     `json:"description"`
	}

	AssignmentData struct {
		Data Task `json:"data"`
	}

	AssignmentArrayData struct {
		Data []Assignment `json:"data"`
	}

	StatusAssignment struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	StatusAssignmentArrayData struct {
		Data []StatusAssignment `json:"data"`
	}
)
