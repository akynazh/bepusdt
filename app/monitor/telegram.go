package monitor

import (
	"github.com/akynazh/bepusdt/app/log"
	"github.com/akynazh/bepusdt/app/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var err error

func BotStart(version string) {
	var BotApi = telegram.GetBotApi()
	if BotApi == nil {

		return
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := BotApi.GetUpdatesChan(u)
	if err != nil {
		log.Error("TG Bot GetUpdatesChan Error:", err)

		return
	}

	//telegram.SendWelcome(version)

	// 监听消息
	for _u := range updates {
		if _u.Message != nil {
			if !_u.FromChat().IsPrivate() {

				continue
			}

			telegram.HandleMessage(_u.Message)
		}
		if _u.CallbackQuery != nil {

			telegram.HandleCallback(_u.CallbackQuery)
		}
	}
}
