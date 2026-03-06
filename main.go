package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

// TODO: обработка команд /start, /help, /example
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	token := os.Getenv("TELEGRAM_APITOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Авторизован на аккаунте %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // Игнорируем не-текстовые обновления
			continue
		}

		// Вызываем нашу логику из logic.go
		result := transcribe(update.Message.Text)

		// Отправляем ответ
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
		bot.Send(msg)
	}
}
