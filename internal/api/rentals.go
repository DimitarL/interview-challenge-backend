package api

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/DimitarL/rental/internal/model"
	"github.com/DimitarL/rental/internal/service"
)

var ErrRepeatedParam = errors.New("parameter is repeated more than once")

const (
	PriceMinParam = "price_min"
	PriceMaxParam = "price_max"
	LimitParam    = "limit"
	OffsetParam   = "offset"
	SortParam     = "sort"
	IDsParam      = "ids"
	NearParam     = "near"

	CoordinatesLen = 2
)

type GetRentalsHandler struct {
	rentalSvc *service.Service
}

func NewGetRentalsHandler(rentalSvc *service.Service) *GetRentalsHandler {
	return &GetRentalsHandler{
		rentalSvc: rentalSvc,
	}
}

func (h *GetRentalsHandler) GetRentals(w http.ResponseWriter, req *http.Request) {
	criteria, err := extractRentFilterParameters(req.URL.Query())
	if err != nil {
		RespondWithErr(w, http.StatusBadRequest, err)
		return
	}

	rentals, err := h.rentalSvc.GetRentals(req.Context(), criteria)
	if err != nil {
		if errors.Is(err, service.ErrMissingRental) {
			RespondWithErr(w, http.StatusNotFound, fmt.Errorf("error getting all rents %w", err))
		}
		RespondWithInternalErr(w, err)
		return
	}

	apiRentals := make([]Rental, 0, len(rentals))
	for _, rental := range rentals {
		apiRental := translateRental(rental)
		apiRentals = append(apiRentals, apiRental)
	}

	RespondWithJson(w, apiRentals, http.StatusOK)
}

func extractRentFilterParameters(values url.Values) (service.SearchCriteria, error) {
	criteria := service.SearchCriteria{}

	if err := parseOptionalInt(values, PriceMinParam, criteria.PriceMin); err != nil {
		return service.SearchCriteria{}, err
	}

	if err := parseOptionalInt(values, PriceMaxParam, criteria.PriceMax); err != nil {
		return service.SearchCriteria{}, err
	}

	if err := parseOptionalInt(values, LimitParam, criteria.Limit); err != nil {
		return service.SearchCriteria{}, err
	}

	if err := parseOptionalInt(values, OffsetParam, criteria.Offset); err != nil {
		return service.SearchCriteria{}, err
	}

	if err := parseSort(values, &criteria); err != nil {
		return service.SearchCriteria{}, err
	}

	if err := parseIDs(values, &criteria); err != nil {
		return service.SearchCriteria{}, err
	}

	if err := parseNear(values, &criteria); err != nil {
		return service.SearchCriteria{}, err
	}

	return criteria, nil
}

func parseOptionalInt(values url.Values, paramName string, dest *int) error {
	if !values.Has(paramName) {
		return nil
	}

	val, err := parseInt(values, paramName)
	if err != nil {
		return err
	}

	dest = &val
	return nil
}

func parseInt(values url.Values, paramName string) (int, error) {
	paramValue := values[paramName]
	if len(paramValue) > 1 {
		return 0, fmt.Errorf("query parameter '%s' must be used only once: %w", paramName, ErrRepeatedParam)
	}

	intVal, err := strconv.Atoi(paramValue[0])
	if err != nil {
		return 0, err
	}

	return intVal, nil
}

func parseSort(values url.Values, criteria *service.SearchCriteria) error {
	if !values.Has(SortParam) {
		return nil
	}

	paramValue := values[SortParam]
	if len(paramValue) > 1 {
		return fmt.Errorf("query parameter '%s' must be used only once: %w", SortParam, ErrRepeatedParam)
	}
	criteria.Sort = &paramValue[0]
	return nil
}

func parseIDs(values url.Values, criteria *service.SearchCriteria) error {
	if !values.Has(IDsParam) {
		return nil
	}
	paramValue := values[IDsParam]
	if len(paramValue) > 1 {
		return fmt.Errorf("query parameter '%s' must be used only once: %w", IDsParam, ErrRepeatedParam)
	}

	criteria.IDs = make([]int, 0, len(paramValue))

	strIDs := strings.Split(paramValue[0], ",")
	for _, strID := range strIDs {
		id, err := strconv.Atoi(strID)
		if err != nil {
			return fmt.Errorf("failed converting ID from string to int: %w", err)
		}
		criteria.IDs = append(criteria.IDs, id)
	}

	return nil
}

func parseNear(values url.Values, criteria *service.SearchCriteria) error {
	if !values.Has(NearParam) {
		return nil
	}

	nearParamPair := values[NearParam]
	if len(nearParamPair) > 1 {
		return fmt.Errorf("query parameter '%s' must be used only once: %w", NearParam, ErrRepeatedParam)
	}

	strCoords := strings.Split(nearParamPair[0], ",")
	if len(strCoords) != CoordinatesLen {
		return fmt.Errorf("query parameter '%s' must contain only %d numbers", NearParam, CoordinatesLen)
	}

	var (
		coords model.Coordinates
		err    error
	)

	coords.Lat, err = strconv.ParseFloat(strCoords[0], 64)
	if err != nil {
		return fmt.Errorf("query parameter '%s' has invalid lat: %w", NearParam, err)
	}

	coords.Lng, err = strconv.ParseFloat(strCoords[1], 64)
	if err != nil {
		return fmt.Errorf("query parameter '%s' has invalid lng: %w", NearParam, err)
	}

	criteria.Near = &coords
	return nil
}
