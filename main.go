package main

import (
	"fmt"

	"github.com/MTplusWebSystem/GoBotKit/botkit"
)


func main() {
	bot := botkit.BotInit{
		Token: "5398155583:AAF9SA5cFDb5LLvYoGkQLjhdTw9JVR6R2tg",
	}
	for {
		if bot.ReceiveData(){
			go func() {
				bot.Handler("callback_query",func(event string) {
					fmt.Println(bot.ChatID)
					fmt.Println("tipo:callback_query(", event,")")
					if event == "!cadastro" {
						bot.SendMessages("Nome: \nSobre nome: \n idade: \n")
					}
				})
			}()
			go func() {
				bot.Handler("commands",func(event string) {
					fmt.Println("tipo:commands(",event,")")
					if event == "/menu" {
						layout := map[string]interface{}{
							"inline_keyboard": [][]map[string]interface{}{
								{
									{"text": "Cadastrar", "callback_data": "!cadastro"},
									{"text": "Sair", "callback_data": "!sair"},
								},
							},
						}
						bot.SendButton("Bem-vindo ao menu",layout)
					}
				})
			}()

			go func() {
				bot.Handler("messages", func(event string) {
					fmt.Println("tipo:messages(",event,")")
					if event == "olá"{
						bot.SendMessages("Olá tudo-bem")
					}
				})
			}()
		}
	}
}







