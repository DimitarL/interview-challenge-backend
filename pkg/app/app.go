package app

import (
	"fmt"
	"net/http"

	"github.com/DimitarL/interview-challenge-backend/pkg/handler"
	"github.com/DimitarL/interview-challenge-backend/pkg/storage"
	"github.com/gorilla/mux"
)

type Application struct {
	RentVehicles *handler.RentHandler
}

func NewApplication() *Application {
	st := storage.NewAppStorage()

	return &Application{RentVehicles: handler.NewRentHandler(st)}
}

func (a *Application) Start(host string, port int) error {
	router := a.createRouter()

	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router)
}

func (a *Application) createRouter() *mux.Router {
	router := mux.NewRouter()

	rentRouter := router.PathPrefix("/rentals").Subrouter()
	rentRouter.HandleFunc("", a.RentVehicles.GetManyHandler).Methods("GET")
	// rentRouter.HandleFunc("/{id}", a.RentVehicles.GetByIdHandler).Methods("GET")

	return router
}
