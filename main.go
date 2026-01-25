package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const TG_BOT_TOKEN = ""
const TG_CHAT_ID = 0

func main() {
	var result int
	randomNumberToGuess := -1

	bot, err := tgbotapi.NewBotAPI(TG_BOT_TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Text == "/math" {
				randomNumber1 := rand.IntN(101)
				randomNumber2 := rand.IntN(101)
				result = randomNumber1 + randomNumber2
				n1 := strconv.Itoa(randomNumber1)
				n2 := strconv.Itoa(randomNumber2)

				msg := tgbotapi.NewMessage(update.Message.From.ID, n1+" + "+n2)
				// msg.ReplyToMessageID = update.Message.MessageID

				m, err := bot.Send(msg)
				fmt.Println(m, err)
			}

			resultStr := strconv.Itoa(result)

			if update.Message.Text == resultStr {
				msg := tgbotapi.NewMessage(TG_CHAT_ID, "You are not stupid!")
				m, err := bot.Send(msg)
				fmt.Println(m, err)
			}

			if update.Message.Text == "/num" {
				randomNumberToGuess = rand.IntN(101)
				fmt.Println("Загадал: ", randomNumberToGuess)
				msg := tgbotapi.NewMessage(update.Message.From.ID, "Угадай число от 0 до 100 ")
				bot.Send(msg)
			}

			randomNumberToGuessAsString := strconv.Itoa(randomNumberToGuess)
			if randomNumberToGuess != -1 {
				if update.Message.Text == randomNumberToGuessAsString {
					msg := tgbotapi.NewMessage(update.Message.From.ID, "Поздравляю! Вы угадали!!!")
					bot.Send(msg)
				} else {
					userTextAsNumber, err := strconv.Atoi(update.Message.Text)
					if err == nil {
						if userTextAsNumber > randomNumberToGuess {
							msg := tgbotapi.NewMessage(update.Message.From.ID, "Загаданное число меньше")
							bot.Send(msg)
						} else {
							msg := tgbotapi.NewMessage(update.Message.From.ID, "Загаданное число больше")
							bot.Send(msg)
						}
					}
				}
			}
		}
	}
}
