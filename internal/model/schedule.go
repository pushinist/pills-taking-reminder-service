package model

import "time"

type Schedule struct {
	ID           int64      `json:"id"`
	UserID       string     `json:"user_id"`
	MedicineName string     `json:"medicine_name"`
	Frequency    Frequency  `json:"frequency"`
	StartDate    time.Time  `json:"start_date"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
}
