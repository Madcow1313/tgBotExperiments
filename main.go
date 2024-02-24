package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("google", "google.com"),
		tgbotapi.NewInlineKeyboardButtonURL("yandex", "ya.ru"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("12356789")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	updateCfg := tgbotapi.NewUpdate(0)
	updateCfg.Timeout = 60
	updateChan := bot.GetUpdatesChan(updateCfg)

	for update := range updateChan {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "sayhi":
					msg.Text = "Hello there, " + update.Message.From.FirstName
				case "status":
					msg.Text = "Everything work as expected"
				default:
					msg.Text = "Unknown command"
				}
			} else {
				switch update.Message.Text {
				case "open":
					msg.ReplyMarkup = numericKeyboard
				}
			}
			if _, err = bot.Send(msg); err != nil {
				log.Panic(err)
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err = bot.Request(callback); err != nil {
				log.Panic(err)
			}
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err = bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}
