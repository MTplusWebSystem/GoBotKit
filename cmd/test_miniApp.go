package main

import (
	"os"

	"github.com/MTplusWebSystem/GoBotKit/botkit"
	"github.com/MTplusWebSystem/GoBotKit/utils"
)

func main() {
	// Tenta ler o diretório ./templates
	files, err := os.ReadDir("./templates")

	// Se der erro (diretório não existe) ou estiver vazio, cria os templates
	if err != nil || len(files) == 0 {
		utils.CreateTemplate()
	} else {
		botkit.MiniAppEX()
	}
}
