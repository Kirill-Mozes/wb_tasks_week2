package model

import "time"

type Event struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

func (e *Event) Validate() bool {
	_, err := time.Parse("2006-01-02", e.Date)
	if err != nil {
		return false
	}
	return e.Name != ""
}
