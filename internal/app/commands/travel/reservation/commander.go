package reservation

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/travel/reservation"
)

type ReservationCommander struct {
	bot                *tgbotapi.BotAPI
	reservationService *service.DummyReservationService
}

func NewReservationCommander(bot *tgbotapi.BotAPI) *ReservationCommander {
	return &ReservationCommander{
		bot:                bot,
		reservationService: service.NewDummyReservationService(),
	}
}

func (c *ReservationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("ReservationCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *ReservationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "delete":
		c.Delete(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "edit":
		c.Edit(msg)
	case "new":
		c.New(msg)
	default:
		//c.Default(msg)
	}
}
