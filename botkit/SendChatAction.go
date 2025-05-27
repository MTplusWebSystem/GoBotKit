package botkit

import (
	"fmt"

	"github.com/MTplusWebSystem/GoBotKit/requests"
)

func (b *BotInit) SendChatAction() error {
	Url := fmt.Sprintf("https://api.telegram.org/bot%s/sendChatAction", b.Token)
    params := map[string]interface{}{
        "chat_id": b.ChatID,
        "action":  "typing", 
    }
    requests.POST(Url, "", params)

	return nil
}