package service

import "github.com/DimitarL/rental/internal/model"

type SearchCriteria struct {
	PriceMin *int
	PriceMax *int
	Limit    *int
	Offset   *int
	IDs      []int
	Near     *model.Coordinates
	Sort     *string
}
