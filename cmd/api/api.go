package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/HimandriSharma/ecommerce/service/cart"
	"github.com/HimandriSharma/ecommerce/service/order"
	"github.com/HimandriSharma/ecommerce/service/products"
	"github.com/HimandriSharma/ecommerce/service/usr"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := usr.NewStore(s.db)
	userHandler := usr.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := products.NewStore(s.db)
	productHandler := products.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)
	log.Println("Listening on:", s.addr)

	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}
