package telebot

import (
	"dayli/dayliModel"
	"dayli/utils"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var botKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Список моих записей"),
	),
)

func StartMessageAndKeyboard(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	startMsg := "Дневник Бот \nНапишите любое сообщение оно будет добавлено в ваш дневник"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, startMsg)
	msg.ReplyMarkup = botKeyboard
	bot.Send(msg)
}

func AddEntryMsgToDb(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	dayliEntry := dayliModel.DayliEntry{Text: update.Message.Text, Id_user: update.Message.Chat.ID}
	dayliEntry.AddEntry()

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Запись сохранена")
	bot.Send(msg)
}

func GetEntrys(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	entrys := dayliModel.GetEntrysByUser(update.Message.Chat.ID)
	textMessage := formatString(entrys)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, textMessage)
	bot.Send(msg)
}

func ProcessingCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	idEntry, err := strconv.Atoi(update.Message.Command())
	utils.CheckError(err)
	entrys := dayliModel.GetEntryByIdEntryAndIdUser(idEntry, update.Message.Chat.ID)

	textMessage := formatString(entrys)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, textMessage)
	msg.ReplyMarkup = getKeybordToEntry(idEntry)
	bot.Send(msg)

}

func DeleteEntry(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	idUser := update.CallbackQuery.Message.Chat.ID
	idEntry, err := strconv.Atoi(update.CallbackQuery.Data[7:])
	utils.CheckError(err)
	dayliModel.DeleteEntry(idEntry, idUser)

	msg := tgbotapi.NewMessage(idUser, "Запись удалена")
	bot.Send(msg)
}

func formatString(entrys []dayliModel.DayliEntry) string {
	var textMessage string
	for _, entry := range entrys {
		textMessage += fmt.Sprintf("/%d - %s\n", entry.Id_entry, entry.Text)
	}
	return textMessage
}

func getKeybordToEntry(idEntry int) tgbotapi.InlineKeyboardMarkup {
	command := fmt.Sprintf("delete-%d", idEntry)
	var entryKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Удалить", command),
		),
	)
	return entryKeyboard
}
