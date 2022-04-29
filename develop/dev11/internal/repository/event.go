package repository

import (
	"dev11/internal/model"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	DAY   int = 1
	WEEK  int = 7
	MONTH int = 30
)

type CalendarCache struct {
	sync.RWMutex
	data map[uuid.UUID]model.Event
}

func NewCalendarCache() *CalendarCache {
	return &CalendarCache{
		RWMutex: sync.RWMutex{},
		data:    make(map[uuid.UUID]model.Event),
	}
}
func (m *CalendarCache) CreateEvent(event model.Event) (uuid.UUID, error) {
	m.Lock()
	defer m.Unlock()
	event.ID = uuid.New()
	m.data[event.ID] = event
	return event.ID, nil
}

func (m *CalendarCache) UpdateEvent(event model.Event) error {
	m.Lock()
	defer m.Unlock()
	_, ok := m.data[event.ID]
	if !ok {
		return model.ErrNotFound
	}
	m.data[event.ID] = event
	return nil
}

func (m *CalendarCache) DeleteEvent(id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()
	_, ok := m.data[id]
	if !ok {
		return model.ErrNotFound
	}
	delete(m.data, id)
	return nil
}

func (m *CalendarCache) GetEvent(interval int, date time.Time) ([]model.Event, error) {
	m.RLock()
	defer m.RUnlock()

	var events []model.Event
	var startDate time.Time
	var endDate time.Time

	switch interval {
	case DAY:
		startDate = date.AddDate(0, 0, -1)
		endDate = date.AddDate(0, 0, 1)
	case WEEK:
		startDate = date.AddDate(0, 0, -4)
		endDate = date.AddDate(0, 0, 4)
	case MONTH:
		startDate = date.AddDate(0, 0, -16)
		endDate = date.AddDate(0, 0, 16)
	}

	for _, v := range m.data {
		if startDate.Before(v.Date) && endDate.After(v.Date) {
			events = append(events, v)
		}
	}
	return events, nil
}
