package main

import (
	"fmt"
	"time"

	"github.com/MTplusWebSystem/GoBotKit/botkit"
)

func main() {
	bot := botkit.BotInit{
		Token: "5398155583:AAF9SA5cFDb5LLvYoGkQLjhdTw9JVR6R2tg",
	}
	for {
		if bot.ReceiveData() {
			go func() {
				bot.Handler("callback_query", func(event string) {
					fmt.Println(bot.ChatID)
					fmt.Println("tipo:callback_query(", event, ")")
					if event == "!cadastro" {
						bot.SendMessages("Nome: \nSobre nome: \n idade: \n")
					}
					if event == "!sair" {
						bot.ReplyToMessage(bot.QueryMessageID, "Tem certeza ?\n")
						layout := map[string]interface{}{
							"inline_keyboard": [][]map[string]interface{}{
								{
									{"text": "Cancelar", "callback_data": "!cancelar"},
									{"text": "Continuar", "callback_data": "!continuar"},
								},
							},
						}
						bot.SendButton("ainda não terminou o cadastro", layout)
					}
				})
			}()
			go func() {
				bot.Handler("commands", func(event string) {
					fmt.Println("tipo:commands(", event, ")")
					if event == "/menu" {
						layout := map[string]interface{}{
							"inline_keyboard": [][]map[string]interface{}{
								{
									{"text": "Cadastrar", "callback_data": "!cadastro"},
									{"text": "Sair", "callback_data": "!sair"},
								},
							},
						}
						bot.SendButton("Bem-vindo ao menu", layout)
					} else if event == "/start" {
						layout := map[string]interface{}{
							"inline_keyboard": [][]map[string]interface{}{
								{
									{"text": "Suporte", "callback_data": "!suporte"},
									{"text": "Painel", "callback_data": "!painel"},
								},
							},
						}
						bot.ReplyToPhotoButton("./boas-vinda.jpg", layout)

					}
				})
			}()

			go func() {
				bot.Handler("messages", func(event string) {
					fmt.Println("tipo:messages(", event, ")")
					if event == "olá" {
						bot.SendMessages("Olá tudo-bem")
					} else if event == "stiker" {
						bot.SendSticker("CAACAgIAAxkBAAIFD2XSt7jvTz70u4qx4tMdF8GG0jFPAALOBQAClvoSBSyTK1YenM2tNAQ")
						for i := 0; i < 2; i++ {
							bot.DeleteMessage(bot.MessageID - i)
						}
					}
				})
			}()
		}
		time.Sleep(1 * time.Second)
	}
}
