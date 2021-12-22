package models

import "time"

type Severity struct {
	ID          int       `json:"id"`
	ServiceName string    `json:"service_name"`
	Severity    string    `json:"severity"`
	Count       int       `json:"count"`
	CreatedAt   time.Time `json:"created_at"`
}
