package botkit

import (
	"fmt"
	"strconv"
	"github.com/MTplusWebSystem/GoBotKit/requests"
)

func (b *BotInit) SendMessages(message string, format ...string) {
	parseMode := "HTML"
	chatID := b.ChatID

	if len(format) > 0 {
		parseMode = format[0]
	}

	if len(format) > 1 {
		id, err := strconv.Atoi(format[1])
		if err != nil {
			fmt.Println("Erro ao converter chatID:", err)
			return
		}
		chatID = id
	}

	params := map[string]interface{}{
		"chat_id":    chatID,
		"parse_mode": parseMode,
		"text":       message,
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Token)
	requests.POST(url, "", params)
}
