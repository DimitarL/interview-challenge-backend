package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/DimitarL/rental/internal/model"
)

var (
	ErrMissingRental = errors.New("rental is missing")
)

type Store interface {
	// GetRental returns a rental by ID if one is found.
	// Else it returns a ErrMissingRental.
	GetRental(context.Context, int) (model.Rental, error)
	GetRentals(context.Context, SearchCriteria) ([]model.Rental, error)
}

type Service struct {
	st Store
}

func NewService(st Store) *Service {
	return &Service{
		st: st,
	}
}

func (s *Service) GetRental(ctx context.Context, rentalID int) (model.Rental, error) {
	rental, err := s.st.GetRental(ctx, rentalID)
	if err != nil {
		if errors.Is(err, ErrMissingRental) {
			return model.Rental{}, fmt.Errorf("missing rental: %w", ErrMissingRental)
		}
		return model.Rental{}, fmt.Errorf("failed getting rent: %w", err)
	}

	return rental, nil
}

func (s *Service) GetRentals(ctx context.Context, criteria SearchCriteria) ([]model.Rental, error) {
	rentals, err := s.st.GetRentals(ctx, criteria)
	if err != nil {
		return []model.Rental{}, fmt.Errorf("failed to get rentals: %w", err)
	}

	return rentals, nil
}
