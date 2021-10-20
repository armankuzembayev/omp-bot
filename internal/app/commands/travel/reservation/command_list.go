package reservation

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ReservationCommander) List(inputMessage *tgbotapi.Message) {
	reservations, err := c.reservationService.List(0, pageLimit)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	outputMessage := "Here all the reservations: \n\n"
	for i, res := range reservations {
		outputMessage += fmt.Sprintf("ID: %v\n", i)
		outputMessage += fmt.Sprintf("User:\n%v\nTicket:\n%v\nExpired:\n %v", res.User, res.Ticket, res.Expired)
		outputMessage += "\n\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessage)

	if c.reservationService.GetLength() > pageLimit {
		serializedData, _ := json.Marshal(
			CallbackListData{
				Cursor: pageLimit,
				Limit:  pageLimit,
			})

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", "travel__reservation__list__"+string(serializedData)),
			),
		)
	}

	c.bot.Send(msg)
}
