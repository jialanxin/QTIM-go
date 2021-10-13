package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func telebotEcho() {
	bot, err := tgbotapi.NewBotAPI("2071123501:AAF9ncd_srDaRJ8UxoEuBz47eAS5T1NSru4")
	if err != nil {
		fmt.Println(err)
	}
	bot.Debug = true
	fmt.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		fmt.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func telebotInformMe() {
	bot, err := tgbotapi.NewBotAPI("2071123501:AAF9ncd_srDaRJ8UxoEuBz47eAS5T1NSru4")
	if err != nil {
		fmt.Println(err)
	}
	bot.Debug = true
	fmt.Printf("Authorized on account %s", bot.Self.UserName)
	msg := tgbotapi.NewMessage(2009774989, "测试")
	bot.Send(msg)
}
