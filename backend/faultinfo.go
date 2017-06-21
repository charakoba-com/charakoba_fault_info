package faultinfo

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server represents an API server
type Server struct {
	*mux.Router
}

// New returns a new API server
func New() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.setupRoutes()
	return s
}

// Run server
func Run(l string) error {
	s := New()
	log.Printf("Server listening on %s", l)

	return http.ListenAndServe(l, s.Router)
}

func (s *Server) setupRoutes() {
	log.Printf("Initialize Routings")

	r := s.Router
	r.HandleFunc(`/`, PostInfoHandler).Methods(`POST`)
	r.HandleFunc(`/`, GetInfoListHandler).Methods(`GET`)
	r.HandleFunc(`/{id:[0-9]+}`, UpdateInfoHandler).Methods(`PUT`)
	r.HandleFunc(`/types`, GetTypesHandler).Methods(`GET`)
	r.HandleFunc(`/services`, GetServicesHandler).Methods(`GET`)
}
