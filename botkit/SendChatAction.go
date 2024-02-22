package botkit

import (
	"fmt"

	"github.com/MTplusWebSystem/GoBotKit/requests"
)

func (b *BotInit) SendChatAction() error {
	UrlAction := fmt.Sprintf("https://api.telegram.org/bot%s/sendChatAction", b.Token)
    paramsAction := map[string]interface{}{
        "chat_id": b.ChatID,
        "action":  "typing", 
    }
    requests.POST(UrlAction, "", paramsAction)

	return nil
}