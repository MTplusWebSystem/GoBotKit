package botkit

import (
	"fmt"

	"github.com/MTplusWebSystem/GoBotKit/requests"
	"github.com/MTplusWebSystem/GoBotKit/system"
)

func (b *BotInit) ReplyToMessage(messageID , text string) error {
	Url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Token)
	var chat_id int = 0
	if b.ChatID == 0{
		chat_id = b.CallbackID
	} else{
		chat_id = b.ChatID
	}
	params := map[string]interface{}{
		"chat_id":           chat_id,
		"text":              text,
		"reply_to_message_id": messageID,
	}

	_, err := requests.POST(Url,"", params)
	system.NilError("erro ao enviar a mensagem de replay:", err)

	return nil
}
