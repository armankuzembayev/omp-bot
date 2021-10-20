package travel

import (
	"fmt"
	"time"
)

type Reservation struct {
	User    *User
	Ticket  *Ticket
	Expired time.Time
}

type ReservationTemp struct {
	User    *User
	Ticket  *Ticket
	Expired string
}

type User struct {
	Name, Surname string
}

type Ticket struct {
	From string
	To   string
	Seat string
}

type Expired struct {
	Expired time.Time
}

func (r *Reservation) String() string {
	return fmt.Sprintf("Reservation:\nUser:\n%v\nTicket:\n%v\nExpired:\n %v", r.User, r.Ticket, r.Expired)
}

func (u *User) String() string {
	return fmt.Sprintf(" Name: %v\n Surname: %v", u.Name, u.Surname)
}

func (t *Ticket) String() string {
	return fmt.Sprintf(" From: %v\n To: %v\n Seat: %v", t.From, t.To, t.Seat)
}
