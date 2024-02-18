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
					fmt.Println("tipo:callback_query(", event,")")
				})
			}()
			go func() {
				bot.Handler("commands",func(event string) {
					fmt.Println("tipo:commands(",event,")")
				})
			}()

			go func() {
				bot.Handler("messages", func(event string) {
					fmt.Println("tipo:messages(",event,")")
				})
			}()
		}
	}
}







