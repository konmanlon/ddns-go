package ddns

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMsgToTgBot(text string) error {
	bot, err := tgbotapi.NewBotAPI(Config.BotToken)
	if err != nil {
		return fmt.Errorf("bot token %s", err)
	}

	text = fmt.Sprintf("\\#DDNS\n```\n%s\n```", text)
	msg := tgbotapi.NewMessage(Config.ChatID, text)
	msg.ParseMode = "MarkdownV2"
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
