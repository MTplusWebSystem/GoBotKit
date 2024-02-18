package botkit

import (
	"fmt"
	"net/url"
	"github.com/MTplusWebSystem/GoBotKit/requests"
	"github.com/MTplusWebSystem/GoBotKit/system"
)


func (b *BotInit) SendMessages(message string){
	Url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s",  b.Token,b.ChatID, url.QueryEscape(message))
	_, err := requests.GET(Url)
	system.NilError("", err)
	return 
}