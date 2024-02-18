package main

import (
	"fmt"

	"github.com/MTplusWebSystem/GoBotKit/botkit"
)


func main() {
	bot := botkit.BotInit{
		Token: "<KEY>",
	}
	for {
		if bot.ReceiveData(){
			fmt.Println("Online")
		}
	}
}







