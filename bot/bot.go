package telebot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(telegramToken string) (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	return bot, updates
}

func LoopUpdates(bot tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.CallbackQuery == nil && update.Message == nil {
			continue
		} else if update.Message != nil {
			messageAnswer(update, &bot)
		} else if update.CallbackQuery != nil {
			callbackQueryAnswer(update, &bot)
		}
	}
}

func messageAnswer(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message.IsCommand() {
		if update.Message.Command() == "start" {
			StartMessageAndKeyboard(update, bot)
		} else {
			ProcessingCommand(update, bot)
		}
	} else if update.Message.Text == "Список моих записей" {
		GetEntrys(update, bot)

	} else if update.Message.Text != "" {
		AddEntryMsgToDb(update, bot)

	}
}

func callbackQueryAnswer(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.CallbackQuery.Data[:7] == "delete-" {
		DeleteEntry(update, bot)
	}
}
