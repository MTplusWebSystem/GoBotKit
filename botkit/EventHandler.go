package botkit

func (b *BotInit) Handler(listen string, callbacks func(event string)) {
	canal_commands := make(chan string)
	canal_messages := make(chan string)
	callback_query := make(chan string)
	go func() {
		if listen == "callback_query" && len(b.CallbackQuery) > 0 && b.CallbackQuery[0] == '!' {
			callback_query <- b.CallbackQuery
		} else if listen == "commands" && len(b.Text) > 0 && b.Text[0] == '/' {
			canal_commands <- b.Text
		} else if listen == "messages" && len(b.Text) > 0 && b.Text[0] != '/' {
			canal_messages <- b.Text
		}
	}()
	select {
	case command := <-canal_commands:
		callbacks(command)
	case message := <-canal_messages:
		callbacks(message)
	case query := <-callback_query:
		callbacks(query)
	}
}