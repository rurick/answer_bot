package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

)

//pageStart  выводит первую страницу по команде start
func pageStart(update *tgbotapi.Update, chatID int64, bot *tgbotapi.BotAPI) {
	name := update.Message.From.FirstName
	text := "<b>Компания 'Рога и Копыта'</b>\n\n" +
		fmt.Sprintf("🤝Здравствуйте, %s!\n", name) +
		"Рога и копыта для нужд гребеночной и мундштучной промышленности. За нашими рогами бегут со всех копыт со всего Мира!\n\n" +
		"<i>Вы попали в хорошую Компанию!😎</i>\n"

	reply := tgbotapi.NewMessage(chatID, "")
	reply.ParseMode = "HTML"
	reply.ReplyMarkup = keyBoardRoot
	reply.Text = text
	bot.Send(reply)
}

func page1(update *tgbotapi.Update, chatID int64, bot *tgbotapi.BotAPI) {
	text := "💼 О компании\n\nОстап Бендер основывает контору «Рога и копыта» для того, чтобы во время расследования «дела А. И. Корейко» «смешаться с бодрой массой служащих». Впервые о подобного рода учреждении, созданном «для нужд гребёночной и пуговичной промышленности», упоминалось в записных книжках Ильфа в 1928 году; позже аналогичный образ появился в рассказе Ильи Арнольдовича «Случай в конторе»[91]. Кроме того, заведения, занимавшиеся заготовкой «когтей и хвостов» и «горчицы и щёлока», фигурировали в повести Ильфа и Петрова «Тысяча и один день, или Новая Шахерезада»"
	reply := tgbotapi.NewMessage(chatID, "")
	reply.ParseMode = "HTML"
	reply.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Хочу узнать побольше...", "https://ru.wikipedia.org/wiki/%D0%97%D0%BE%D0%BB%D0%BE%D1%82%D0%BE%D0%B9_%D1%82%D0%B5%D0%BB%D1%91%D0%BD%D0%BE%D0%BA#%C2%AB%D0%A0%D0%BE%D0%B3%D0%B0_%D0%B8_%D0%BA%D0%BE%D0%BF%D1%8B%D1%82%D0%B0%C2%BB"),
		),
	)
	reply.Text = text
	bot.Send(reply)
}

func page2(update *tgbotapi.Update, chatID int64, bot *tgbotapi.BotAPI) {
	text := "<b>📚 Наша продукция</b>\n\n Медведь, балалайка и шапка-ушанка символизируют дремучую Россию, заснеженную и непонятную. Но это скорее киношные стереотипы, чем реальное представление о самой большой стране на Земле. Сегодня нас представляют совсем другие бренды, известные во всем мире российским происхождением."
	reply := tgbotapi.NewMessage(chatID, "")
	reply.ParseMode = "HTML"
	reply.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🦏 Рога", "inline2.1"),
			tgbotapi.NewInlineKeyboardButtonData("🐎 Копыта", "inline2.2"),
		),
	)
	reply.Text = text
	bot.Send(reply)
}

func page2Edit(update *tgbotapi.Update, chatID int64, bot *tgbotapi.BotAPI) {
	text := "<b>📚 Наша продукция</b>\n\n Медведь, балалайка и шапка-ушанка символизируют дремучую Россию, заснеженную и непонятную. Но это скорее киношные стереотипы, чем реальное представление о самой большой стране на Земле. Сегодня нас представляют совсем другие бренды, известные во всем мире российским происхождением."
	reply := tgbotapi.NewEditMessageText(chatID, update.CallbackQuery.Message.MessageID, text)
	reply.ParseMode = "HTML"
	bot.Send(reply)
	kb := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🦏 Рога", "inline2.1"),
			tgbotapi.NewInlineKeyboardButtonData("🐎 Копыта", "inline2.2"),
		),
	)
	reply2 := tgbotapi.NewEditMessageReplyMarkup(chatID, update.CallbackQuery.Message.MessageID, kb)
	bot.Send(reply2)
}

func page2_1(update *tgbotapi.Update, chatID int64, bot *tgbotapi.BotAPI) {
	text := "🦏 Рога\n\nРог — образование на головах у парнокопытных млекопитающих из семейств полорогих, вилороговых, оленевых и жирафовых, а также у непарнокопытных семейства носороговых. Кроме того, рога имели многие вымершие млекопитающие из других семейств и отрядов: диноцераты, протоцератиды, бронтотерии и другие. В широком смысле этого слова рогами могут называться и внешне схожие образования на теле у других животных, например у лепидозавров (трёхрогие хамелеоны), архозавров (птицы-носороги, трицератопсы) и жуков. \n\n"
	text += "<a href='https://upload.wikimedia.org/wikipedia/commons/thumb/3/31/Capra_nubiana%2C_Israel_1.jpg/330px-Capra_nubiana%2C_Israel_1.jpg'>&#8205;</a>"
	reply := tgbotapi.NewEditMessageText(chatID, update.CallbackQuery.Message.MessageID, "")
	reply.ParseMode = "HTML"
	reply.Text = text
	bot.Send(reply)
	kb := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⏪ Назад", "inline2.back1"),
		),
	)
	reply2 := tgbotapi.NewEditMessageReplyMarkup(chatID, update.CallbackQuery.Message.MessageID, kb)
	bot.Send(reply2)
}

func page2_2(update *tgbotapi.Update, chatID int64, bot *tgbotapi.BotAPI) {
	text := "Копы́то — твёрдое роговое образование вокруг дистальных пальцевых фаланг копытных млекопитающих. Для парнокопытных принят термин копытце[1]. В анатомическом отношении копыта соответствуют ногтям у человека. Копыто является модифицированной кожей, у которой отсутствует нижний слой, а эпидермис превращён в мозоль. Патологии копыт изучает ветеринарная ортопедия. "
	text += "<a href='https://upload.wikimedia.org/wikipedia/commons/thumb/3/33/Hufschnitt.jpg/375px-Hufschnitt.jpg'>&#8205;</a>"
	reply := tgbotapi.NewMessage(chatID, "")
	reply.ParseMode = "HTML"
	reply.Text = text
	reply.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⏪ Назад", "inline2.back2"),
		),
	)
	bot.Send(reply)
}

func page3(update *tgbotapi.Update, chatID int64, bot *tgbotapi.BotAPI) {
	text := "📞 Контакты\n\nБоты любой сложности. Пишите @YuriyIovkov"
	reply := tgbotapi.NewMessage(chatID, "")
	reply.ParseMode = "HTML"
	reply.Text = text
	bot.Send(reply)
}
