package store

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/DimitarL/rental/internal/model"
	"github.com/DimitarL/rental/internal/service"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	pool *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{
		pool: pool,
	}
}

func (s *Store) GetRental(ctx context.Context, rentalID int) (model.Rental, error) {
	query := fmt.Sprintf(`SELECT %s
		FROM rentals JOIN users ON rentals.user_id = users.id
		WHERE rentals.id = $1;`,
		rentalFields(),
	)

	row := s.pool.QueryRow(ctx, query, rentalID)

	rental := model.Rental{}
	err := row.Scan(expandRentalEntry(&rental)...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Rental{}, service.ErrMissingRental
		}
		return model.Rental{}, fmt.Errorf("failed scanning row: %w", err)
	}

	return rental, nil
}

func (s *Store) GetRentals(ctx context.Context, criteria service.SearchCriteria) ([]model.Rental, error) {
	query := fmt.Sprintf(`SELECT %s
		FROM rentals JOIN users ON rentals.user_id = users.id %s`,
		rentalFields(), buildRentalsFilter(criteria),
	)

	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query rentals: %w", err)
	}
	defer rows.Close()

	var rentals []model.Rental

	for rows.Next() {
		rental := model.Rental{}

		err := rows.Scan(expandRentalEntry(&rental)...)
		if err != nil {
			return nil, fmt.Errorf("failed scanning for a rental entry: %w", err)
		}

		rentals = append(rentals, rental)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed iterating rows: %w", err)
	}

	return rentals, nil
}

func rentalFields() string {
	return `rentals.id, rentals.name, rentals.description, rentals.type, rentals.vehicle_make, rentals.vehicle_model,
			rentals.vehicle_year, rentals.vehicle_length, rentals.sleeps, rentals.primary_image_url,
			rentals.price_per_day, rentals.home_city, rentals.home_state, rentals.home_zip, rentals.home_country,
			rentals.lat, rentals.lng, users.id, users.first_name, users.last_name`
}

func expandRentalEntry(rental *model.Rental) []interface{} {
	return []interface{}{
		&rental.ID, &rental.Name, &rental.Description, &rental.Type, &rental.Make, &rental.Model,
		&rental.Year, &rental.Length, &rental.Sleeps, &rental.PrimaryImageURL, &rental.PricePerDay.Day, &rental.Location.City,
		&rental.Location.State, &rental.Location.Zip, &rental.Location.Country, &rental.Location.Coordinates.Lat,
		&rental.Location.Coordinates.Lng, &rental.User.ID, &rental.User.FirstName, &rental.User.LastName,
	}
}

func buildRentalsFilter(criteria service.SearchCriteria) string {
	var restrictions []string

	if criteria.PriceMin != nil {
		restrictions = append(restrictions, fmt.Sprintf("price_per_day >= %d", *criteria.PriceMin))
	}

	if criteria.PriceMax != nil {
		restrictions = append(restrictions, fmt.Sprintf("price_per_day <= %d", *criteria.PriceMax))
	}

	if len(criteria.IDs) > 0 {
		var strIDs []string
		for _, id := range criteria.IDs {
			strIDs = append(strIDs, strconv.Itoa(id))
		}

		restrictions = append(restrictions, fmt.Sprintf("rentals.id IN (%s)", strings.Join(strIDs, ", ")))
	}

	if criteria.Near != nil {
		restrictions = append(restrictions, fmt.Sprintf(`
			acos(
				sin(radians(%f)) * sin(radians(lat))
				+ cos(radians(%f)) * cos(radians(lat)) * cos(radians(%f)
				- radians(lng))
			) * 6371 <= 100`, criteria.Near.Lat, criteria.Near.Lat, criteria.Near.Lng))
	}

	builder := strings.Builder{}
	if len(restrictions) > 0 {
		builder.WriteString(" WHERE " + strings.Join(restrictions, " AND "))
	}

	if criteria.Sort != nil {
		checkSortCriteriaType(criteria, &builder)
	}

	if criteria.Limit != nil {
		builder.WriteString(fmt.Sprintf(" LIMIT %d", *criteria.Limit))
	}

	if criteria.Offset != nil {
		builder.WriteString(fmt.Sprintf(" OFFSET %d", *criteria.Offset))
	}

	return builder.String()
}

func checkSortCriteriaType(criteria service.SearchCriteria, builder *strings.Builder) {
	if *criteria.Sort == "price" {
		builder.WriteString(" ORDER BY price_per_day")
	}
	if *criteria.Sort == "year" {
		builder.WriteString(" ORDER BY vehicle_year")
	}
	if *criteria.Sort == "name" {
		builder.WriteString(" ORDER BY name")
	}
	if *criteria.Sort == "make" {
		builder.WriteString(" ORDER BY vehicle_make")
	}
}
