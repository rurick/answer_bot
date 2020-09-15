package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const (
	//Button1 -
	Button1 = "💼 О компании"
	//Button2 -
	Button2 = "📚 Наша продукция"
	//Button3 -
	Button3 = "📞 Контакты"
)

//keyBoard - Клавиатура
var keyBoardRoot = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(Button1),
		tgbotapi.NewKeyboardButton(Button2),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(Button3),
	),
)
