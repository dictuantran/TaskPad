package models

import "time"

type TaskMetric struct {
	Day    time.Time `json:"day"`
	Effort float32   `json:"effort"`
}
