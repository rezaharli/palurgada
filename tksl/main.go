package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1248386508:AAHTEibsHh7MFt58DK6YUUQbGbA_XLnFvpg")
	if err != nil {
		fmt.Println("error:", err.Error())
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		fmt.Println("error:", err.Error())
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if update.Message.Command() == "name" {
				sendName(bot, update)
			}
		} else if update.Message.Document != nil {
			fmt.Println("")
			fmt.Println("Dokumen nih")
			fmt.Printf("[%s] %s", update.Message.From.UserName, update.Message.Document.FileID)
			getFile(bot, update)
		}
	}
}

func sendName(bot *(tgbotapi.BotAPI), update tgbotapi.Update) (tgbotapi.Message, error) {
	return sendMessage(bot, update.Message.Chat.ID, "My name is tksl")
}

func getFile(bot *(tgbotapi.BotAPI), update tgbotapi.Update) (tgbotapi.Message, error) {
	fileConf := tgbotapi.FileConfig{
		FileID: update.Message.Document.FileID,
	}

	fileInfo, err := bot.GetFile(fileConf)
	if err != nil {
		fmt.Println("error:", err.Error())
	}

	fmt.Println("-----")
	fmt.Println(fileInfo.Link("1248386508:AAHTEibsHh7MFt58DK6YUUQbGbA_XLnFvpg"))

	err = DownloadFile("logo.xlsx", fileInfo.Link("1248386508:AAHTEibsHh7MFt58DK6YUUQbGbA_XLnFvpg"))
	if err != nil {
		panic(err)
	}

	return sendMessage(bot, update.Message.Chat.ID, "Data file uploaded")
}

func sendMessage(bot *(tgbotapi.BotAPI), chatID int64, message string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, message)

	return bot.Send(msg)
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
