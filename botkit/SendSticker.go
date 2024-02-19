package botkit

import (
	"fmt"
	"github.com/MTplusWebSystem/GoBotKit/requests"
)

func (b *BotInit) SendSticker(stickerFileID string) {
	params := map[string]interface{}{
		"chat_id": b.ChatID,
		"sticker": stickerFileID,
	}
	Url := fmt.Sprintf("https://api.telegram.org/bot%s/sendSticker", b.Token)
	requests.POST(Url, "application/json", params)
}
