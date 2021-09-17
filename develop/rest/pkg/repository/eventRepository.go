package repository

import (
	"database/sql"
	"fmt"
	"simplerest/pkg/model"
	"time"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (e *EventRepository) GetEventsForCountDays(today, tomorrow string) ([]model.Event, error) {
	stmt, err := e.db.Prepare("" +
		"SELECT * " +
		"FROM events " +
		"WHERE date BETWEEN $1 AND $2")
	if err != nil {
		fmt.Println("Error Prepare in GetEventsForCountDay")
		return nil, err
	}

	rows, err := stmt.Query(today, tomorrow)
	if err != nil {
		fmt.Println("Error Query in GetEventsForCountDay")
		return nil, err
	}
	defer rows.Close()

	events := []model.Event{}

	for rows.Next() {
		type EventDB struct {
			ID          int64
			Name        string
			Description string
			Date        time.Time
		}

		var eventDb EventDB
		err = rows.Scan(&eventDb.ID, &eventDb.Name, &eventDb.Description, &eventDb.Date)
		if err != nil {
			fmt.Println("Error in GetEventsForCountDay")
			return nil, err
		}

		event := model.Event{
			ID:          eventDb.ID,
			Name:        eventDb.Name,
			Description: eventDb.Description,
			Date:        eventDb.Date.Format("2006-01-02"),
		}
		events = append(events, event)
	}

	return events, nil
}

func (e *EventRepository) CreateEvent(event model.Event) (int64, error) {
	query := `
		INSERT INTO events
		(name, description, date)
		VALUES ($1, $2, $3)
		RETURNING id`

	var id int64
	row := e.db.QueryRow(query, event.Name, event.Description, event.Date)
	if err := row.Scan(&id); err != nil {
		fmt.Println("Error in CreateEvent")
		return 0, err
	}

	return id, nil
}

func (e *EventRepository) UpdateEvent(event model.Event) (model.Event, error) {
	query := `
		UPDATE events
		SET name=$1, description=$2, date=$3
		WHERE id=$4
		RETURNING id, name, description, date`

	type EventDB struct {
		ID          int64
		Name        string
		Description string
		Date        time.Time
	}

	var eventDb EventDB
	row := e.db.QueryRow(query, event.Name, event.Description, event.Date, event.ID)
	if err := row.Scan(&eventDb.ID, &eventDb.Name, &eventDb.Description, &eventDb.Date); err != nil {
		fmt.Println("Error in UpdateEvent")
		return model.Event{}, err
	}

	updatedEvent := model.Event{
		ID:          eventDb.ID,
		Name:        eventDb.Name,
		Description: eventDb.Description,
		Date:        eventDb.Date.Format("02-01-2006"),
	}

	return updatedEvent, nil
}

func (e *EventRepository) DeleteEvent(id int64) error {
	query := `
			DELETE FROM events
			WHERE id = $1`

	if _, err := e.db.Exec(query, id); err != nil {
		fmt.Println("Error in DeleteEvent")
		return err
	}

	return nil
}
