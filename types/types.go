package types

import "time"

type Task struct {
	Id              int       `json:"id"`
	TaskDescription string    `json:"taskDescpription"`
	Status          string    `json:"status"` //todo, in progress and done
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
