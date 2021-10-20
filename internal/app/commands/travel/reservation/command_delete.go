package reservation

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ReservationCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	reservationId, err := strconv.Atoi(args)
	if err != nil || reservationId < 0 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Reservation id {%v} is invalid", args))
		c.bot.Send(msg)
		return
	}

	isRemoved, err := c.reservationService.Remove(uint64(reservationId))
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	if isRemoved {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "User was deleted")
		c.bot.Send(msg)
	}

}
