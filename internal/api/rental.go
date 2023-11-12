package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/DimitarL/rental/internal/service"
	"github.com/gorilla/mux"
)

type GetRentalHandler struct {
	rentalSvc *service.Service
}

func NewGetRentalHandler(rentalSvc *service.Service) *GetRentalHandler {
	return &GetRentalHandler{
		rentalSvc: rentalSvc,
	}
}

func (h *GetRentalHandler) GetRental(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	rentalID, err := strconv.Atoi(params["id"])
	if err != nil {
		err = fmt.Errorf("parameter 'id': %w", err)
		RespondWithErr(w, http.StatusBadRequest, err)
		return
	}

	rental, err := h.rentalSvc.GetRental(req.Context(), rentalID)
	if err != nil {
		if errors.Is(err, service.ErrMissingRental) {
			RespondWithErr(w, http.StatusNotFound, fmt.Errorf("rental of vehicle with id %d not found", rentalID))
		}
		RespondWithInternalErr(w, err)
		return
	}

	apiRental := translateRental(rental)
	RespondWithJson(w, apiRental, http.StatusOK)
}
