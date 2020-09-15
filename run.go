package main

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//Run - новый поток обработчика
func Run(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var (
		chatID int64
	)
	if update.Message != nil {
		chatID = update.Message.Chat.ID
	}
	if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.Message.Chat.ID
	}

	if update.Message != nil {
		//Проверка команд
		if update.Message.Command() == "start" {
			pageStart(update, chatID, bot)
		} else {
			switch update.Message.Text {
			case Button1:
				page1(update, chatID, bot)
			case Button2:
				page2(update, chatID, bot)
			case Button3:
				page3(update, chatID, bot)
			default:
				c := tgbotapi.NewDeleteMessage(chatID, update.Message.MessageID)
				c1 := tgbotapi.NewMessage(chatID, "⛔️ Ничего не знаю про это...")
				m, _ := bot.Send(c1)
				time.Sleep(2 * time.Second)
				bot.Send(c)
				c = tgbotapi.NewDeleteMessage(chatID, m.MessageID)
				bot.Send(c)
			}
		}
	}
	if update.CallbackQuery != nil {
		switch update.CallbackQuery.Data {
		case "inline2.1":
			page2_1(update, chatID, bot)
		case "inline2.2":
			page2_2(update, chatID, bot)
		case "inline2.back1":
			page2Edit(update, chatID, bot)
		case "inline2.back2":
			m := tgbotapi.NewDeleteMessage(chatID, update.CallbackQuery.Message.MessageID)
			bot.Send(m)
		}
	}

}
