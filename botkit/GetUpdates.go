package botkit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MTplusWebSystem/GoBotKit/requests"
	"github.com/MTplusWebSystem/GoBotKit/system"
)

/*
Criando as funçoes para criação de bot

GetUpdates

ReceiveData

SendMessage

DeleteMessage

EditMessage -> fica para proxima atualização

SendPhoto

SendAudio -> fica para proxima atualização

SendDocument -> fica para proxima atualização

SendSticker

SendVideo -> fica para proxima atualização

KeybordButton

InlineKeyboardButton

InlineKeyboardMarkup

ReplayButton

ReplayToMessage

ReplayToPhotoButton
*/

type BotInit struct {
	Token string
	UpdateID	int
	MessageID    int
	ChatID       int
	CallbackID   int
	Username     string
	Text         string
	CallbackQuery string
	QueryMessageID int
	ID           int
}

type Update struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"message"`
	CallbackQuery struct {
		Data    string `json:"data"`
		Message struct {
			MessageID int `json:"message_id"`
			Chat struct {
				ID int `json:"id"`
			} `json:"chat"`
		}
	} `json:"callback_query"`
}
var Start int = 0



func (b *BotInit) GetUpdates() ([]Update, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?timeout=100&offset=%d", b.Token, b.UpdateID+1)

	response, err := requests.GET(url)
	system.NilError("",err)
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status de resposta HTTP inválido: %s", response.Status)
	}
	
	responseBody, err := ioutil.ReadAll(response.Body)
	system.NilError("Erro ao ler resposta HTTP:", err)
	
	var updates struct {
		OK     bool     `json:"ok"`
		Result []Update `json:"result"`
	}
	if err := json.Unmarshal(responseBody, &updates); err != nil {
		fmt.Println("Erro ao decodificar resposta JSON:", err)
		return nil, fmt.Errorf("Erro ao decodificar resposta JSON: %s", err)
	}

	if !updates.OK {
		return nil, fmt.Errorf("Erro na solicitação de atualizações")
	}
	
	for _, update := range updates.Result {
		b.ChatID = update.Message.Chat.ID
	}
	return updates.Result, nil
}