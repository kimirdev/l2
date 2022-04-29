package repository

import (
	"dev11/internal/model"
	"time"

	"github.com/google/uuid"
)

type Calendar interface {
	CreateEvent(event model.Event) (uuid.UUID, error)
	UpdateEvent(event model.Event) error
	DeleteEvent(id uuid.UUID) error
	GetEvent(interval int, date time.Time) ([]model.Event, error)
}

type Repository struct {
	Calendar
}
