package reservation

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/travel"
	"strconv"
	"strings"
	"time"
)

const (
	layoutDate = "2006-01-02 15:04"
)

func (c *ReservationCommander) Edit(inputMessage *tgbotapi.Message) {
	splitArgs := strings.SplitN(inputMessage.CommandArguments(), ", ", 2)

	if len(splitArgs) != 2 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Input is not correct"))
		c.bot.Send(msg)
		return
	}

	reservationId, err := strconv.Atoi(splitArgs[0])
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Reservation id is not valid"))
		c.bot.Send(msg)
		return
	}

	reservationTemp := travel.ReservationTemp{}

	if err = json.Unmarshal([]byte(splitArgs[1]), &reservationTemp); err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Reservation structure is not valid: %v", splitArgs[1]))
		c.bot.Send(msg)
		return
	}

	reservation := travel.Reservation{}
	if reservationTemp.User == nil || reservationTemp.Ticket == nil || reservationTemp.Expired == "" {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Information about reservation was not typed correctly"))
		c.bot.Send(msg)
		return
	}

	reservation.User = reservationTemp.User
	reservation.Ticket = reservationTemp.Ticket
	reservation.Expired, err = time.Parse(layoutDate, reservationTemp.Expired)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	if err = c.reservationService.Update(uint64(reservationId), reservation); err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Edited successfully"))
	c.bot.Send(msg)
}
