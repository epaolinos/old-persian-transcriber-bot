// Package main provides the Telegram bot entry point and
// command handling logic for the Old Persian translator.
package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

// main is the entry point of the application. It initializes the bot,
// sets up configuration, and starts the update polling loop.
func main() {
	// Load environment variables from the .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	// Retrieve the Telegram API token from environment variables.
	token := os.Getenv("TELEGRAM_APITOKEN")

	// Initialize the bot with the provided token.
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// Enable debug mode to see more detailed logs about incoming updates.
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Set up the update configuration with a long-polling timeout.
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Start receiving updates from Telegram.
	updates := bot.GetUpdatesChan(u)

	// Iterate through each incoming update.
	for update := range updates {
		// Ignore updates that do not contain a message.
		if update.Message == nil {
			continue
		}

		// Handle bot commands (messages starting with '/').
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				text := "Hello! I can transcribe your text into the Old Persian cuneiform.\n\n" +
					"If you want to try it out, click /example, copy the phrase from the Behistun Inscription and send it back to me.\n\n" +
					"For more details, click /help.\n\n" +
					"Made by @epaolinos (Pavel Egizaryan)"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				bot.Send(msg)

			case "example":
				exampleText := "θātiy Dārayavauš xšāyaθiya vašnā Ahuramazdāha adam xšāyaθiya amiy Ahuramazdāha xšaçam manā frābara"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, exampleText)
				bot.Send(msg)

			case "help":
				helpText := "How to use:\n\n" +
					"1. Send a Latin transcription (e.g., `da`).\n" +
					"2. You can use apostrophes for special characters (`a'` → `ā`, `s'` → `š`, `t'` → `θ`, `c'` → `ç`).\n" +
					"3. Numbers are converted automatically, though signs for 200 and above remain uncertain.\n\n" +
					"I will suggest logograms if detected."
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpText)
				msg.ParseMode = "Markdown"
				bot.Send(msg)

			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know this command.")
				bot.Send(msg)
			}

			continue
		}

		// Process regular text messages: transcribe and look for logograms.
		result := transcribe(update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
		bot.Send(msg)
	}
}
