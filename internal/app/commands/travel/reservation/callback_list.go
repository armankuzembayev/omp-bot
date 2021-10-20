package reservation

import (
	"encoding/json"
	"fmt"
	"log"
	"math"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const (
	pageLimit = 5
)

type CallbackListData struct {
	Cursor uint64 `json:"cursor"`
	Limit  uint64 `json:"limit"`
}

func (c *ReservationCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("ReservationCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	outputMessage := ""

	reservations, _ := c.reservationService.List(parsedData.Cursor, parsedData.Limit)
	for i, res := range reservations {
		outputMessage += fmt.Sprintf("ID: %v\n", uint64(i)+parsedData.Cursor)
		outputMessage += fmt.Sprintf("User:\n%v\nTicket:\n%v\nExpired:\n %v", res.User, res.Ticket, res.Expired) + "\n\n"
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMessage)
	buttons := getInlineKeyboardButtons(parsedData, c.reservationService.GetLength())
	if len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(buttons...),
		)
	}
	c.bot.Send(msg)
}

func getInlineKeyboardButtons(parsedData CallbackListData, reservationsSize uint64) []tgbotapi.InlineKeyboardButton {
	inlineKeyboardButtons := make([]tgbotapi.InlineKeyboardButton, 0, 2)

	if parsedData.Cursor > 0 {
		cursor := uint64(math.Max(0, float64(parsedData.Cursor-pageLimit)))

		data, _ := json.Marshal(
			&CallbackListData{
				Cursor: cursor,
				Limit:  pageLimit,
			})

		inlineKeyboardButtons = append(inlineKeyboardButtons, tgbotapi.NewInlineKeyboardButtonData("Prev page", "travel__reservation__list__"+string(data)))
	}

	cursor := parsedData.Cursor + parsedData.Limit
	if reservationsSize > cursor {
		data, _ := json.Marshal(
			&CallbackListData{
				Cursor: cursor,
				Limit:  pageLimit,
			})

		inlineKeyboardButtons = append(inlineKeyboardButtons, tgbotapi.NewInlineKeyboardButtonData("Next page", "travel__reservation__list__"+string(data)))
	}

	return inlineKeyboardButtons
}
