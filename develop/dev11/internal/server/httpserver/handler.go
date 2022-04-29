package httpserver

import (
	"dev11/internal/model"
	"dev11/internal/repository"
	"net/http"
	"time"
)

const (
	TimeFormat = "2006-01-02"
)

func (s *Server) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusBadRequest, "Wrong method")
		return
	}

	var e model.EventDto

	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := e.Validate(); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	uid, err := s.Service.CalendarService.CreateEvent(e)
	if err != nil {
		errorResponse(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	resultResponse(w, http.StatusCreated, uid)
}

func (s *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusBadRequest, "Wrong method")
		return
	}

	var e model.Event

	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.Service.CalendarService.DeleteEvent(e.ID); err != nil {
		errorResponse(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	resultResponse(w, http.StatusOK, "event deleted")
}

func (s *Server) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusBadRequest, "Wrong method")
		return
	}

	var e model.Event

	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.Service.CalendarService.UpdateEvent(e); err != nil {
		errorResponse(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	resultResponse(w, http.StatusOK, "event updated")
}

func (s *Server) eventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, http.StatusBadRequest, "Wrong method")
		return
	}

	date, err := time.Parse(TimeFormat, r.URL.Query().Get("date"))

	if err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	events, err := s.Service.CalendarService.GetEvent(repository.DAY, date)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	resultResponse(w, http.StatusOK, events)

}
func (s *Server) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, http.StatusBadRequest, "Wrong method")
		return
	}

	date, err := time.Parse(TimeFormat, r.URL.Query().Get("date"))

	if err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	events, err := s.Service.CalendarService.GetEvent(repository.WEEK, date)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	resultResponse(w, http.StatusOK, events)
}
func (s *Server) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, http.StatusBadRequest, "Wrong method")
		return
	}

	date, err := time.Parse(TimeFormat, r.URL.Query().Get("date"))

	if err != nil {
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	events, err := s.Service.CalendarService.GetEvent(repository.MONTH, date)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	resultResponse(w, http.StatusOK, events)
}
