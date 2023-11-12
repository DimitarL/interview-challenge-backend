package api

import "github.com/DimitarL/rental/internal/model"

type Rental struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Type            string   `json:"type"`
	Make            string   `json:"make"`
	Model           string   `json:"model"`
	Year            int      `json:"year"`
	Length          float64  `json:"length"`
	Sleeps          int      `json:"sleeps"`
	PrimaryImageURL string   `json:"primary_image_url"`
	PricePerDay     Price    `json:"price_per_day"`
	Location        Location `json:"location"`
	User            User     `json:"user"`
}

type Price struct {
	Day int `json:"day"`
}

type Location struct {
	City    string  `json:"city"`
	State   string  `json:"state"`
	Zip     string  `json:"zip"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func translateRental(rental model.Rental) Rental {
	return Rental{
		ID:              rental.ID,
		Name:            rental.Name,
		Description:     rental.Description,
		Type:            rental.Type,
		Make:            rental.Make,
		Model:           rental.Model,
		Year:            rental.Year,
		Length:          rental.Length,
		Sleeps:          rental.Sleeps,
		PrimaryImageURL: rental.PrimaryImageURL,
		PricePerDay: Price{
			Day: rental.PricePerDay.Day,
		},
		Location: Location{
			City:    rental.Location.City,
			State:   rental.Location.State,
			Zip:     rental.Location.Zip,
			Country: rental.Location.Country,
			Lat:     rental.Location.Coordinates.Lat,
			Lng:     rental.Location.Coordinates.Lng,
		},
		User: User{
			ID:        rental.User.ID,
			FirstName: rental.User.FirstName,
			LastName:  rental.User.LastName,
		},
	}
}
