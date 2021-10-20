package travel

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/travel/reservation"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type ReservationCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type TravelCommander struct {
	bot             *tgbotapi.BotAPI
	travelCommander Commander
}

func NewTravelCommander(bot *tgbotapi.BotAPI) *TravelCommander {
	return &TravelCommander{
		bot:             bot,
		travelCommander: reservation.NewReservationCommander(bot),
	}
}

func (c *TravelCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "reservation":
		c.travelCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("TravelCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *TravelCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "reservation":
		c.travelCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("TravelCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
