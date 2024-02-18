package botkit

import (
	"encoding/json"
	"fmt"

	"github.com/MTplusWebSystem/GoBotKit/requests"
	"github.com/MTplusWebSystem/GoBotKit/system"
)

func (bot *BotInit) SendButton(menu string, layout interface{}) {
    send := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", bot.Token)
    
    var replyMarkup map[string]interface{}
    replyMarkupBytes, err := json.Marshal(layout)
	system.NilError("Erro ao converter layout para JSON:", err)
    err = json.Unmarshal(replyMarkupBytes, &replyMarkup)
    system.NilError("Erro ao converter JSON para mapa:", err)

    params := map[string]interface{}{
        "chat_id":      bot.ChatID,
        "text":         menu,
        "parse_mode":   "HTML",
        "reply_markup": replyMarkup,
    }
    requests.POST(send, "", params)
}