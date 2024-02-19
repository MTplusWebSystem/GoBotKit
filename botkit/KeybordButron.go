package botkit

import "github.com/MTplusWebSystem/GoBotKit/requests"

func (b *BotInit) KeyboardButton(text string, keyboard interface{}) {
    reqBody := map[string]interface{}{
        "chat_id":      b.ChatID,
        "text":         text,
        "reply_markup": keyboard,
    }
    resp, _ := requests.POST("https://api.telegram.org/bot"+b.Token+"/sendMessage", "application/json", reqBody)
    defer resp.Body.Close()
}
