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
		if event == "!suporte" {
			bot.ForceReplyToMessage(bot.QueryMessageID ,"Nome do usúario")
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

					if event == "/submenu" {
						layout := map[string]interface{}{
							"keyboard": [][]map[string]interface{}{
								{
									{"text": "Suporte"},
									{"text": "Painel","callback_data": "!sair"},
								},
							},
							"resize_keyboard":   true,
							"one_time_keyboard": true,
						}						
                        bot.KeyboardButton("Clique no botão ao lado esquerdo do emojii ", layout)
                    }else if event == "/menu" {
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

					} else if event == "/FORCE"{
						layout := map[string]interface{}{
                            "inline_keyboard": [][]map[string]interface{}{
                                {
                                    {"text": "Suporte", "callback_data": "!suporte"},
                                    {"text": "Painel", "callback_data": "!painel"},
                                },
                            },
                        }
                        bot.SendButton("Menu", layout)
					}
				})
			}()

			go func() {
				if bot.ReplyMessageText == "Nome do usúario"{
					fmt.Println("O nome de usúario foi",bot.Text)
				}
				bot.Handler("messages", func(event string) {
					switch event {
						case "negrito":
							bot.SendMessages("```go \n fmt.Println(O nome de usúario foi,bot.Text)```")
					}
					fmt.Println("tipo:messages(", event, ")")
					if event == "projeto" {
						bot.SendDocument("./GoBotKit-main.zip", "projeto GoBotKit", "zip")
					} else if event == "stiker" {
						bot.SendSticker("CAACAgIAAxkBAAIFD2XSt7jvTz70u4qx4tMdF8GG0jFPAALOBQAClvoSBSyTK1YenM2tNAQ")
						time.Sleep(1 * time.Second)
						bot.DeleteMessage(bot.MessageID + 1)
					}
				})
			}()
		}
		time.Sleep(1 * time.Second)
	}
}
