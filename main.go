package main

import (
	"fmt"
	"github.com/MTplusWebSystem/GoBotKit/system"
)

func main() {
	data := string(system.Scan("./arquivo.txt"))
	fmt.Println(data)

	resultado := system.Regex("[0-9]+", "texto123")
	fmt.Println(resultado)

	dados := map[string]interface{}{"nome": "Exemplo", "idade": 30}
	err := system.WriteJSON("dados.json", dados)
	system.NilError("Erro ao escrever arquivo JSON:", err)
	
	aleatorio := system.Random("0-9", false)
	fmt.Println(aleatorio)

	chave := system.KeyGenerator(4)
	fmt.Println(chave)
}