package botkit

import (
	"fmt"

	"github.com/MTplusWebSystem/GoBotKit/requests"
	"github.com/MTplusWebSystem/GoBotKit/utils"
)

func (b *BotInit) ReplyToMessage(messageID int , text string) error {
	Url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Token)
	params := map[string]interface{}{
		"chat_id":           b.ChatID,
		"text":              text,
		"reply_to_message_id": messageID,
	}

	_, err := requests.POST(Url,"", params)
	utils.NilError("erro ao enviar a mensagem de replay:", err)

	return nil
}

