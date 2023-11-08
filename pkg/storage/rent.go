package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/DimitarL/interview-challenge-backend/common"
	"github.com/DimitarL/interview-challenge-backend/pkg/entities"

	pgx "github.com/jackc/pgx/v4"
)

const (
	getByIdRentQuery = `SELECT rentals.id, rentals.name, rentals.description, rentals.type, rentals.vehicle_make,
	rentals.vehicle_model, rentals.vehicle_year, rentals.vehicle_length, rentals.sleeps, rentals.primary_image_url,
	rentals.price_per_day, rentals.home_city, rentals.home_state, rentals.home_zip, rentals.home_country, rentals.lat,
	rentals.lng, users.id, users.first_name, users.last_name
	FROM rentals JOIN users ON rentals.user_id = users.id
	WHERE rentals.id = $1;`
	listAllSelectPartOfQuery = `SELECT rentals.id, rentals.name, rentals.description, rentals.type, rentals.vehicle_make,
	rentals.vehicle_model, rentals.vehicle_year, rentals.vehicle_length, rentals.sleeps, rentals.primary_image_url,
	rentals.price_per_day, rentals.home_city, rentals.home_state, rentals.home_zip, rentals.home_country, rentals.lat,
	rentals.lng, users.id, users.first_name, users.last_name
	FROM rentals JOIN users ON rentals.user_id = users.id`
)

func (a AppStorage) ListAllRentVehicles(params *common.RentalQueryParameters, addWherePart bool) ([]entities.RentResponse, error) {
	rents := []entities.RentResponse{}
	var getAllQuery string

	if addWherePart {
		getAllQuery = listAllSelectPartOfQuery + buildWhereClause(params)
	} else {
		getAllQuery = listAllSelectPartOfQuery
	}

	rows, err := a.conn.Query(context.Background(), getAllQuery)
	if err != nil {
		return rents, err
	}
	defer rows.Close()

	for rows.Next() {
		var rentResponseData entities.RentResponse

		err := rows.Scan(expandRentalEntry(&rentResponseData)...)
		if err != nil {
			return rents, fmt.Errorf("failed to scan row: %w", err)
		}

		rents = append(rents, rentResponseData)
	}

	if err := rows.Err(); err != nil {
		return rents, err
	}

	return rents, nil
}

func (a AppStorage) GetRentVehicleById(id int) (*entities.RentResponse, error) {
	var rentResponseData entities.RentResponse

	err := a.conn.QueryRow(context.Background(), getByIdRentQuery, id).Scan(expandRentalEntry(&rentResponseData)...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
	}

	return &rentResponseData, nil
}

func buildWhereClause(params *common.RentalQueryParameters) string {
	return newWhereClauseBuilder().AddParamsIfPresent(params)
}

func expandRentalEntry(rent *entities.RentResponse) []interface{} {
	return []interface{}{&rent.ID, &rent.Name, &rent.Description, &rent.Type, &rent.Make, &rent.Model,
		&rent.Year, &rent.Length, &rent.Sleeps, &rent.PrimaryImageURL, &rent.PricePerDay.Day, &rent.Location.City,
		&rent.Location.State, &rent.Location.Zip, &rent.Location.Country, &rent.Location.Lat, &rent.Location.Lng,
		&rent.User.ID, &rent.User.FirstName, &rent.User.LastName}
}
