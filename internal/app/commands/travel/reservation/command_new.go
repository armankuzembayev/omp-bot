package reservation

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/travel"
)

func (c *ReservationCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	reservationTemp := travel.ReservationTemp{}
	reservation := travel.Reservation{}
	if args[0] == '{' {
		err := json.Unmarshal([]byte(args), &reservationTemp)
		if err != nil {
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Reservation structure is not valid: %v", args))
			c.bot.Send(msg)
			return
		}

		if reservationTemp.User == nil || reservationTemp.Ticket == nil || reservationTemp.Expired == "" {
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Information about reservation was not typed correctly"))
			c.bot.Send(msg)
			return
		}

		reservation.User = reservationTemp.User
		reservation.Ticket = reservationTemp.Ticket
		expired, err := time.Parse(layoutDate, reservationTemp.Expired)
		if err != nil {
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
			c.bot.Send(msg)
			return
		}
		reservation.Expired = expired

	} else {
		splittedArgs := strings.Split(args, " ")
		if len(splittedArgs) != 7 {
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Reservation structure is not valid: %v", args))
			c.bot.Send(msg)
			return
		}
		user := &travel.User{
			Name:    splittedArgs[0],
			Surname: splittedArgs[1],
		}
		ticket := &travel.Ticket{
			From: splittedArgs[2],
			To:   splittedArgs[3],
			Seat: splittedArgs[4],
		}

		reservation.User = user
		reservation.Ticket = ticket
		expired, err := time.Parse(layoutDate, splittedArgs[5]+" "+splittedArgs[6])
		if err != nil {
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
			c.bot.Send(msg)
			return
		}
		reservation.Expired = expired
	}

	reservationId, err := c.reservationService.Create(reservation)
	reservationId--
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("New user with id {%v} was created successfully", reservationId),
	)
	c.bot.Send(msg)
}
