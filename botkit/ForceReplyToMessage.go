package botkit

import (
    "fmt"
    "github.com/MTplusWebSystem/GoBotKit/requests"
)

func (b *BotInit) ForceReplyToMessage(messageID int, text string) error {
    Url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Token)
    params := map[string]interface{}{
        "chat_id":            b.ChatID,
        "text":               text,
        "reply_to_message_id": messageID,
        "reply_markup": map[string]interface{}{
            "force_reply": true,
        },
    }
    requests.POST(Url, "", params)
    return nil
}