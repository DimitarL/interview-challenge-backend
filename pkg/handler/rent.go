package handler

import (
	"fmt"
	"net/http"
	"net/url"

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
	params, err := extractRentFilterParameters(r.URL.Query())
	if err != nil {
		common.RespondWithErr(w, http.StatusBadRequest, err)
		return
	}

	addWherePart := !(len(r.URL.Query()) == 0)

	rents, err := rnt.st.ListAllRentVehicles(params, addWherePart)
	if err != nil {
		err := fmt.Errorf("error getting all rents %w", err)
		common.RespondWithInternalErr(w, err)
		return
	}

	common.RespondWithJson(w, rents, http.StatusOK)
}

func extractRentFilterParameters(query url.Values) (*common.RentalQueryParameters, error) {
	return common.ExtractQueryParameters(query)
}
