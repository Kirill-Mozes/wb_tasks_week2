package repository

import (
	"database/sql"
	"simplerest/pkg/model"
)

type Event interface {
	GetEventsForCountDays(today, tomorrow string) ([]model.Event, error)
	CreateEvent(event model.Event) (int64, error)
	UpdateEvent(event model.Event) (model.Event, error)
	DeleteEvent(id int64) error
}

type Repository struct {
	Event
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Event: NewEventRepository(db),
	}
}
