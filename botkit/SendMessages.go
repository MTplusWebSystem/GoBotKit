package botkit

import (
	"fmt"
	"net/url"
	"github.com/MTplusWebSystem/GoBotKit/requests"
	"github.com/MTplusWebSystem/GoBotKit/system"
)


func (b *BotInit) SendMessages(message string){
	var chat_id int = 0
	if b.ChatID == 0 {
		chat_id = b.CallbackID
	} else {
		chat_id = b.ChatID
	}
	Url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s",  b.Token,chat_id, url.QueryEscape(message))
	_, err := requests.GET(Url)
	system.NilError("", err)
	return 
}