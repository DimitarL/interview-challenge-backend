package app

import (
	"fmt"
	"net/http"

	"github.com/DimitarL/interview-challenge-backend/pkg/storage"
	"github.com/gorilla/mux"
)

type Application struct {
}

func NewApplication() *Application {
	storage.NewAppStorage()

	return &Application{}
}

func (a *Application) Start(host string, port int) error {
	router := a.createRouter()

	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router)
}

func (a *Application) createRouter() *mux.Router {
	router := mux.NewRouter()

	return router
}
