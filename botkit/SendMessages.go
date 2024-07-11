package botkit

import (
	"fmt"
	"github.com/MTplusWebSystem/GoBotKit/requests"
)

func (b *BotInit) SendMessages(message string, format ...string) {
    parseMode := "HTML"
    chatID := b.ChatID

    if len(format) > 0 {
        parseMode = format[0]
    }
    
    if len(format) > 1 {
        chatID = format[1]
    }

    params := map[string]interface{}{
        "chat_id":    chatID,
        "parse_mode": parseMode,
        "text":       message,
    }

    Url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Token)
    requests.POST(Url, "", params)
    return
}
