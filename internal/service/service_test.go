package service_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/DimitarL/rental/internal/model"
	"github.com/DimitarL/rental/internal/service"
	"github.com/DimitarL/rental/internal/service/servicefakes"
)

var (
	errGetRental  = errors.New("failed getting rental")
	errGetRentals = errors.New("failed getting rentals")
)

var _ = Describe("Service", func() {
	Context("created", func() {
		It("returns", func() {
			Expect(service.NewService(nil)).NotTo(BeNil())
		})
	})

	Describe("instance", func() {
		var (
			ctx       context.Context
			svc       *service.Service
			fakeSt    *servicefakes.FakeStore
			errAction error
		)

		ItSucceeds := func() {
			GinkgoHelper()
			It("succeeds", func() {
				Expect(errAction).NotTo(HaveOccurred())
			})
		}
		ItFailsWith := func(expected error) {
			GinkgoHelper()
			It("fails", func() {
				Expect(errAction).To(MatchError(expected))
			})
		}

		BeforeEach(func() {
			ctx = context.Background()
			fakeSt = new(servicefakes.FakeStore)
			svc = service.NewService(fakeSt)
		})

		Context("getting rental", func() {
			var rentalID1 = 123
			var rental model.Rental

			JustBeforeEach(func() {
				rental, errAction = svc.GetRental(ctx, rentalID1)
			})

			Context("and rental is found", func() {
				var rental1 = model.Rental{
					ID:   rentalID1,
					Name: "blabla",
				}
				BeforeEach(func() {
					fakeSt.GetRentalReturns(rental1, nil)
				})

				ItSucceeds()
				It("returns rental", func() {
					Expect(rental).To(Equal(rental1))
				})
				It("searches for the correct rental", func() {
					Expect(fakeSt.GetRentalCallCount()).To(Equal(1))

					_, rentalID := fakeSt.GetRentalArgsForCall(0)
					Expect(rentalID).To(Equal(rentalID1))
				})
			})
			Context("and rental is not found", func() {
				BeforeEach(func() {
					fakeSt.GetRentalReturns(model.Rental{}, service.ErrMissingRental)
				})

				ItFailsWith(service.ErrMissingRental)
				It("returns no rental", func() {
					Expect(rental).To(BeZero())
				})
			})
			Context("and rental fails with another error", func() {
				BeforeEach(func() {
					fakeSt.GetRentalReturns(model.Rental{}, errGetRental)
				})
				ItFailsWith(errGetRental)
				It("returns no rental", func() {
					Expect(rental).To(BeZero())
				})
			})
		})
		Context("getting rentals", func() {
			var (
				priceMin  = 5
				limit     = 10
				criteria1 = service.SearchCriteria{
					PriceMin: &priceMin,
					Limit:    &limit,
				}
			)
			var rentals []model.Rental

			JustBeforeEach(func() {
				rentals, errAction = svc.GetRentals(ctx, criteria1)
			})

			Context("and store finds rentals", func() {
				var (
					rental1 = model.Rental{
						ID:   1,
						Name: "foo",
					}
					rental2 = model.Rental{
						ID:   2,
						Name: "bar",
					}
				)
				BeforeEach(func() {
					fakeSt.GetRentalsReturns([]model.Rental{rental1, rental2}, nil)
				})

				ItSucceeds()
				It("returns the rentals", func() {
					Expect(rentals).To(Equal([]model.Rental{rental1, rental2}))
				})
				It("calls the store with the correct criteria", func() {
					Expect(fakeSt.GetRentalsCallCount()).To(Equal(1))

					_, criteria := fakeSt.GetRentalsArgsForCall(0)
					Expect(criteria).To(Equal(criteria1))
				})
			})
			Context("and store fails to find the rentals", func() {
				BeforeEach(func() {
					fakeSt.GetRentalsReturns(nil, errGetRentals)
				})

				ItFailsWith(errGetRentals)
				It("returns no rentals", func() {
					Expect(rentals).To(BeEmpty())
				})
			})
		})
	})
})
