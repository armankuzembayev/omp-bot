package reservation

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ReservationCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"/help__travel__reservation - print list of commands\n"+
			"/get__travel__reservation - get a reservation by id (id starts with 0)\n"+
			"/list__travel__reservation - get a list of reservations\n"+
			"/delete__travel__reservation - delete an existing reservation by id (id starts with 0)\n"+
			"/new__travel__reservation - create a new reservation\n\n"+

			"Reservation in JSON format or space-separated list of arguments\n"+
			"User, Ticket and Expired date must be specified\n"+
			"Expired date must be in the following format: {year-month-day hour:minute}\n\n"+

			"Example: /new__travel__reservation {\"User\": {\"Name\":\"Marcus\", \"Surname\":\"Rashford\"}, \"Ticket\": {\"From\":\"London\", \"To\":\"Manchester\", \"Seat\":\"10B\"}, \"Expired\": \"2021-11-20 07:35\"}\n"+
			"Example: /new__travel__reservation Marcus Rashford London Manchester 10B 2021-11-20 07:35\n\n"+

			"/edit__travel__reservation - edit a reservation by id (id starts with 0)\n\n"+
			"Input format: {reservation id, reservation in JSON format}\n"+
			"User, Ticket and Expired date must be specified\n"+
			"Example: /edit__travel__reservation 0, {\"User\": {\"Name\":\"Marcus\", \"Surname\":\"Rashford\"}, \"Ticket\": {\"From\":\"London\", \"To\":\"Manchester\", \"Seat\":\"10B\"}, \"Expired\": \"2021-11-20 07:35\"}\n",
	)

	c.bot.Send(msg)
}
