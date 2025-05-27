package botkit

import (
	"encoding/json"
	"fmt"

	"github.com/MTplusWebSystem/GoBotKit/requests"
	"github.com/MTplusWebSystem/GoBotKit/utils"
)

func (bot *BotInit) SendButton(menu string, layout interface{},format... string) {
    send := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", bot.Token)
    parseMode := "HTML"
    if len(format) > 0 {
        parseMode = format[0]
    }
    var replyMarkup map[string]interface{}
    replyMarkupBytes, err := json.Marshal(layout)
	utils.NilError("Erro ao converter layout para JSON:", err)
    err = json.Unmarshal(replyMarkupBytes, &replyMarkup)
    utils.NilError("Erro ao converter JSON para mapa:", err)

    params := map[string]interface{}{
        "chat_id":      bot.ChatID,
        "parse_mode":   parseMode,
        "text":         menu,
        "reply_markup": replyMarkup,
    }
    requests.POST(send, "", params)
}