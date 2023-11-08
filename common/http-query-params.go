package common

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type RentalQueryParameters struct {
	PriceMin *int      `json:"price_min,omitempty"`
	PriceMax *int      `json:"price_max,omitempty"`
	Limit    *int      `json:"limit,omitempty"`
	Offset   *int      `json:"offset,omitempty"`
	Ids      string    `json:"ids,omitempty"`
	Near     []float64 `json:"near,omitempty"`
	Sort     string    `json:"sort,omitempty"`
}

func ExtractQueryParameters(query url.Values) (*RentalQueryParameters, error) {
	params := &RentalQueryParameters{}

	if query.Has("price_min") {
		paramValue := query["price_min"]
		if len(paramValue) > 1 {
			return params, fmt.Errorf("query parameter '%s' must be used only once", "price_min")
		}

		priceMin, err := strconv.Atoi(paramValue[0])
		if err != nil {
			return params, err
		} else {
			params.PriceMin = &priceMin
		}
	}
	if query.Has("price_max") {
		paramValue := query["price_max"]
		if len(paramValue) > 1 {
			return params, fmt.Errorf("query parameter '%s' must be used only once", "price_max")
		}

		priceMax, err := strconv.Atoi(paramValue[0])
		if err != nil {
			return params, err
		} else {
			params.PriceMax = &priceMax
		}
	}
	if query.Has("limit") {
		paramValue := query["limit"]
		if len(paramValue) > 1 {
			return params, fmt.Errorf("query parameter '%s' must be used only once", "limit")
		}

		limit, err := strconv.Atoi(paramValue[0])
		if err != nil {
			return params, err
		} else {
			params.Limit = &limit
		}
	}
	if query.Has("offset") {
		paramValue := query["offset"]
		if len(paramValue) > 1 {
			return params, fmt.Errorf("query parameter '%s' must be used only once", "offset")
		}

		offset, err := strconv.Atoi(paramValue[0])
		if err != nil {
			return params, err
		} else {
			params.Offset = &offset
		}
	}
	if query.Has("ids") {
		paramValue := query["ids"]
		if len(paramValue) > 1 {
			return params, fmt.Errorf("query parameter '%s' must be used only once", "ids")
		}
		params.Ids = paramValue[0]
	}
	if query.Has("near") {
		nearParamPair := query["near"]
		if len(nearParamPair) > 1 {
			return params, fmt.Errorf("query parameter '%s' must be used only once", "near")
		}

		nearParamPairStringValues := strings.Split(nearParamPair[0], ",")
		if len(nearParamPairStringValues) != 2 {
			return params, fmt.Errorf("query parameter '%s' must contain only 2 numbers", "near")
		}

		var nearParamPairIntegerValues []float64

		currValue, err := strconv.ParseFloat(nearParamPairStringValues[0], 64)
		if err != nil {
			return params, err
		}
		nearParamPairIntegerValues = append(nearParamPairIntegerValues, currValue)

		currValue, err = strconv.ParseFloat(nearParamPairStringValues[1], 64)
		if err != nil {
			return params, err
		}
		nearParamPairIntegerValues = append(nearParamPairIntegerValues, currValue)

		params.Near = nearParamPairIntegerValues
	}
	if query.Has("sort") {
		paramValue := query["sort"]
		if len(paramValue) > 1 {
			return params, fmt.Errorf("query parameter '%s' must be used only once", "sort")
		}
		params.Sort = paramValue[0]
	}

	return params, nil
}
