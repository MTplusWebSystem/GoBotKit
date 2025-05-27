package botkit

import (
	"fmt"
	"net/http"

	"github.com/MTplusWebSystem/GoBotKit/requests"
	"github.com/MTplusWebSystem/GoBotKit/utils"
)

func (b *BotInit) DeleteMessage(messageID int) error {
	Url := fmt.Sprintf("https://api.telegram.org/bot%s/deleteMessage", b.Token)
	params := map[string]interface{}{
		"chat_id":    b.ChatID,
		"message_id": messageID,
	}

	resp, err := requests.POST(Url, "application/json", params)
	utils.NilError("", err)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Erro ao excluir a mensagem. Código de status: %d", resp.StatusCode)
	}

	return nil
}
