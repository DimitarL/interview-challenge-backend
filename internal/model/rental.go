package model

type Rental struct {
	ID              int
	Name            string
	Description     string
	Type            string
	Make            string
	Model           string
	Year            int
	Length          float64
	Sleeps          int
	PrimaryImageURL string
	PricePerDay     Price
	Location        Location
	User            User
}

type Price struct {
	Day int
}

type Location struct {
	City        string
	State       string
	Zip         string
	Country     string
	Coordinates Coordinates
}

type Coordinates struct {
	Lat float64
	Lng float64
}

type User struct {
	ID        int
	FirstName string
	LastName  string
}
