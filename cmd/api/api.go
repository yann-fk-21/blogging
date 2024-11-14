package api

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	//subRouter := router.PathPrefix("/api/v1")

	fmt.Println("server running...")

	return http.ListenAndServe(s.addr, router)
}
