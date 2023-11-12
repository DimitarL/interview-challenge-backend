package api

import "github.com/gorilla/mux"

type RouterBuilder struct {
	getRentalHandler  *GetRentalHandler
	getRentalsHandler *GetRentalsHandler
}

func NewRouterBuilder(
	getRentalHandler *GetRentalHandler,
	getRentalsHandler *GetRentalsHandler,
) *RouterBuilder {
	return &RouterBuilder{
		getRentalHandler:  getRentalHandler,
		getRentalsHandler: getRentalsHandler,
	}
}

func (b *RouterBuilder) Build() *mux.Router {
	router := mux.NewRouter()

	rentalsRouter := router.PathPrefix("/rentals").Subrouter()
	rentalsRouter.HandleFunc("", b.getRentalsHandler.GetRentals).Methods("GET")
	rentalsRouter.HandleFunc("/{id}", b.getRentalHandler.GetRental).Methods("GET")

	return router
}
