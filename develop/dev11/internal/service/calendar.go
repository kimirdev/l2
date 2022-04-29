package service

import (
	"dev11/internal/model"
	"dev11/internal/repository"
	"time"

	"github.com/google/uuid"
)

type Calendar struct {
	db repository.Calendar
}

func NewCalendar(db repository.Calendar) *Calendar {
	return &Calendar{db: db}
}

func (c *Calendar) CreateEvent(event model.EventDto) (uuid.UUID, error) {
	var e model.Event
	e.Date = time.Time(event.Date)
	e.Name = event.Name
	return c.db.CreateEvent(e)
}
func (c *Calendar) UpdateEvent(event model.Event) error {
	return c.db.UpdateEvent(event)
}
func (c *Calendar) DeleteEvent(id uuid.UUID) error {
	return c.db.DeleteEvent(id)
}
func (c *Calendar) GetEvent(period int, date time.Time) ([]model.Event, error) {
	return c.db.GetEvent(period, date)
}
