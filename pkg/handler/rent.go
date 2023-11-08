package handler

import (
	"fmt"
	"net/http"

	"github.com/DimitarL/interview-challenge-backend/common"
	"github.com/DimitarL/interview-challenge-backend/pkg/storage"
)

type RentHandler struct {
	st *storage.AppStorage
}

func NewRentHandler(st *storage.AppStorage) *RentHandler {
	return &RentHandler{st: st}
}

func (rnt *RentHandler) GetManyHandler(w http.ResponseWriter, r *http.Request) {
	rents, err := rnt.st.ListAllRentVehicles()
	if err != nil {
		err := fmt.Errorf("error getting all rents %w", err)
		common.RespondWithInternalErr(w, err)
		return
	}

	common.RespondWithJson(w, rents, http.StatusOK)
}
