package botkit

import (
	"fmt"

	"github.com/MTplusWebSystem/GoBotKit/requests"
)

func (b *BotInit) ForceReply(message string) error {
    Url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Token)
    params := map[string]interface{}{
        "chat_id":            b.ChatID,
		"text":          message,
        "reply_markup": map[string]interface{}{
            "force_reply": true,
        },
    }
    requests.POST(Url, "", params)

    return nil
}