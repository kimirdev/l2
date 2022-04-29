package service

import (
	"dev11/internal/model"
	"dev11/internal/repository"
	"time"

	"github.com/google/uuid"
)

type CalendarService interface {
	CreateEvent(event model.EventDto) (uuid.UUID, error)
	UpdateEvent(event model.Event) error
	DeleteEvent(id uuid.UUID) error
	GetEvent(period int, date time.Time) ([]model.Event, error)
}

type Service struct {
	CalendarService
}

func NewService(db repository.Calendar) *Service {
	return &Service{
		CalendarService: NewCalendar(db),
	}
}
