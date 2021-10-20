package reservation

import (
	"fmt"
	"math"
	"time"

	"github.com/ozonmp/omp-bot/internal/model/travel"
)

type ReservationService interface {
	Describe(reservationID uint64) (*travel.Reservation, error)
	List(cursor uint64, limit uint64) ([]travel.Reservation, error)
	Create(reservation travel.Reservation) (uint64, error)
	Update(reservationID uint64, reservation travel.Reservation) error
	Remove(reservationID uint64) (bool, error)
}

type DummyReservationService struct{}

func NewDummyReservationService() *DummyReservationService {
	return &DummyReservationService{}
}

func (s *DummyReservationService) GetLength() uint64 {
	return uint64(len(allReservations))
}

func (s *DummyReservationService) isValidId(reservationID uint64) error {
	if reservationID < 0 || reservationID >= s.GetLength() {
		return fmt.Errorf("Reservation id {%d} is out of scope ", reservationID)
	}
	return nil
}

func (s *DummyReservationService) isReservationCorrect(reservation travel.Reservation) error {
	if &reservation.User == nil || &reservation.Ticket == nil || &reservation.Expired == nil {
		return fmt.Errorf("Information about reservation was not typed correctly ")
	}
	return nil
}

func (s *DummyReservationService) Describe(reservationID uint64) (*travel.Reservation, error) {
	if err := s.isValidId(reservationID); err != nil {
		return nil, err
	}
	return &allReservations[reservationID], nil
}

func (s *DummyReservationService) List(cursor uint64, limit uint64) ([]travel.Reservation, error) {
	length := s.GetLength()
	if cursor >= length || cursor < 0 {
		return []travel.Reservation{}, fmt.Errorf("Not valid cursor ")
	}
	minIdx := uint64(math.Min(float64(length), float64(cursor+limit)))
	return allReservations[cursor:minIdx], nil
}

func (s *DummyReservationService) Create(reservation travel.Reservation) (uint64, error) {
	if err := s.isReservationCorrect(reservation); err != nil {
		return 0, err
	}
	allReservations = append(allReservations, reservation)
	return s.GetLength(), nil
}

func (s *DummyReservationService) Update(reservationID uint64, reservation travel.Reservation) error {
	if err := s.isValidId(reservationID); err != nil {
		return err
	}
	if err := s.isReservationCorrect(reservation); err != nil {
		return err
	}
	allReservations[reservationID] = reservation
	return nil
}

func (s *DummyReservationService) Remove(reservationID uint64) (bool, error) {
	if err := s.isValidId(reservationID); err != nil {
		return false, err
	}
	allReservations = append(allReservations[:reservationID], allReservations[reservationID+1:]...)
	return true, nil
}

var allReservations = []travel.Reservation{
	{
		User: &travel.User{
			Name:    "Cristiano",
			Surname: "Ronaldo",
		},
		Ticket: &travel.Ticket{
			From: "Torino",
			To:   "Manchester",
			Seat: "7A",
		},
		Expired: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	},
	{
		User: &travel.User{
			Name:    "Bruno",
			Surname: "Fernandes",
		},
		Ticket: &travel.Ticket{
			From: "Lisbon",
			To:   "Manchester",
			Seat: "18B",
		},
		Expired: time.Date(2019, time.September, 14, 21, 0, 0, 0, time.UTC),
	},
	{
		User: &travel.User{
			Name:    "David",
			Surname: "De Gea",
		},
		Ticket: &travel.Ticket{
			From: "Madrid",
			To:   "Manchester",
			Seat: "1C",
		},
		Expired: time.Date(2015, time.February, 2, 4, 0, 0, 0, time.UTC),
	},
	{
		User: &travel.User{
			Name:    "Paul",
			Surname: "Pogba",
		},
		Ticket: &travel.Ticket{
			From: "Torino",
			To:   "Manchester",
			Seat: "6d",
		},
		Expired: time.Date(2016, time.July, 21, 14, 0, 0, 0, time.UTC),
	},
	{
		User: &travel.User{
			Name:    "Edinson",
			Surname: "Cavani",
		},
		Ticket: &travel.Ticket{
			From: "Paris",
			To:   "Manchester",
			Seat: "21e",
		},
		Expired: time.Date(2019, time.March, 7, 18, 0, 0, 0, time.UTC),
	},
	{
		User: &travel.User{
			Name:    "Mason",
			Surname: "Greenwood",
		},
		Ticket: &travel.Ticket{
			From: "Bradford",
			To:   "Manchester",
			Seat: "11f",
		},
		Expired: time.Date(2014, time.December, 11, 20, 0, 0, 0, time.UTC),
	},
}
