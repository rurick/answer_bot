package main

//answerCallbackQuery
//git clone https://github.com/go-telegram-bot-api/telegram-bot-api.git tgbotapi
//git checkout 99b74b8efaa519636cf7f56afed97b65ecafb512
// GetFileDirectURL

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//Config - конфиг
type Config struct {
	TelegramBotToken string
	DebugLog         bool
}

var configuration Config

//инициализация приложения
func init() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	log.Println("Запуск telegram бота (для выхода нажмите Ctrl+C)...")

	//создаем контексты
	ctx, cansel := context.WithCancel(context.Background()) //контекст программы для обработки событий
	go handleSignals(cansel)

	// TELEGRAM BOT
	// используя токен создаем новый инстанс бота
	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = false
	log.Println("Авторизация в Telegram акаунтом @" + bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, _ := bot.GetUpdatesChan(u)
	log.Println("OK")

	//основной цикл чтения каналов
	for {
		select {
		case <-ctx.Done(): //ждем завершения программы по ctrl+c
			log.Println("Остановка Бота...")
			return

		case update := <-updates: //ждем сообщение из телеги
			go Run(&update, bot)
		}
	}
}

//обработчик сигналов
func handleSignals(cansel context.CancelFunc) {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	for {
		sig := <-sigCh
		switch sig {
		case os.Interrupt:
			cansel()
			return
		}
	}
}
