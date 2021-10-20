package reservation

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ReservationCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	reservationId, err := strconv.Atoi(args)
	if err != nil || reservationId < 0 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Reservation id {%v} is invalid", args))
		c.bot.Send(msg)
		return
	}

	reservation, err := c.reservationService.Describe(uint64(reservationId))
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, reservation.String())
	c.bot.Send(msg)
}
