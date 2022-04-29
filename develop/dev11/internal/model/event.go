package model

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

func (e *Event) Decode(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&e)
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Validate() error {
	if e.Name == "" {
		return fmt.Errorf("name is invalid")
	}

	return nil
}

type EventDto struct {
	Name string     `json:"name"`
	Date customTime `json:"date"`
}

type customTime time.Time

func (j *customTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = customTime(t)
	return nil
}

func (j customTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

func (e *EventDto) Decode(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&e)
	if err != nil {
		return err
	}

	return nil
}

func (e *EventDto) Validate() error {
	if e.Name == "" {
		return fmt.Errorf("name is invalid")
	}

	return nil
}
