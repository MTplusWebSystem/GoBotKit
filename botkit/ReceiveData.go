package botkit

import "fmt"

func (b *BotInit) ReceiveData() bool {
	updates, err := b.GetUpdates()
	if err != nil {
		fmt.Println("Erro ao receber dados:", err)
		return false
	}

	for _, update := range updates {
		b.ID = update.Message.From.ID
		if b.ChatID == 0 {
			b.ChatID = update.CallbackQuery.Message.Chat.ID
		} else {
			b.ChatID = update.Message.Chat.ID
		}
		b.CallbackID = update.CallbackQuery.Message.Chat.ID
		if Start == 0 {
			Start++
		} else if Start != 0 {
			if update.UpdateID != b.UpdateID {
				b.Username = update.Message.From.Username
				b.UpdateID = update.UpdateID
				b.MessageID = update.Message.MessageID 
				b.Text = update.Message.Text
				b.CallbackQuery = update.CallbackQuery.Data
				b.QueryMessageID = update.CallbackQuery.Message.MessageID
				return true
			}
		}

	}

	return false
}