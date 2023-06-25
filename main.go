package main

import (
	telebot "dayli/bot"
	"dayli/dayliModel"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dayliModel.InitDB("./foo.sqlite")
	dayliModel.InitCreateTable()

	bot, updates := telebot.StartBot("123") // Enter Telegram Bot Token
	telebot.LoopUpdates(*bot, updates)
}
