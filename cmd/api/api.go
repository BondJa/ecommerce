package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/HimandriSharma/ecommerce/service/usr"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer{
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error{
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := usr.NewStore(s.db)
	userHandler := usr.NewHandler(userStore)
	userHandler.RegisterRouter(subrouter)

	log.Println("Listening on:", s.addr)

	return http.ListenAndServe(s.addr, router)
}