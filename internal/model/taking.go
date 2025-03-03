package model

import "time"

type Taking struct {
	MedicineName string    `json:"medicine_name"`
	ScheduleID   string    `json:"schedule_id"`
	Time         time.Time `json:"time"`
}
