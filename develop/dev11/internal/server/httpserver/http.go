package httpserver

import (
	"dev11/internal/server/middleware"
	"dev11/internal/service"
	"net"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	Service    *service.Service
}

func New(serv *service.Service) *Server {
	server := &Server{
		httpServer: &http.Server{
			Addr: net.JoinHostPort("", "8081"),
		},
		Service: serv,
	}

	return server
}
func (s *Server) initRouts() {
	http.Handle("/create_event", middleware.Log(http.HandlerFunc(s.createEvent)))
	http.Handle("/update_event", middleware.Log(http.HandlerFunc(s.updateEvent)))
	http.Handle("/delete_event", middleware.Log(http.HandlerFunc(s.deleteEvent)))

	http.Handle("/events_for_day", middleware.Log(http.HandlerFunc(s.eventsForDay)))
	http.Handle("/events_for_week", middleware.Log(http.HandlerFunc(s.eventsForWeek)))
	http.Handle("/events_for_month", middleware.Log(http.HandlerFunc(s.eventsForMonth)))
}

func (s *Server) Run() error {
	s.initRouts()
	return s.httpServer.ListenAndServe()
}
