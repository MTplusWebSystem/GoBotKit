package system

import "fmt"

func NilError(msg string , err error) {
	if err != nil {
		fmt.Println(msg,err)
	}
}

/*
funções para serem criadas

NilError: simplificar as mensagens de erros e diminuir linhas de código

Scan: para leitura simples de arquivos

Regex: Uma forma simples de aplicar regex

WriteJSON: para escrita de arquivos JSON

Random: para geração de valores aleatórios

KeyGenerator: para geração de chaves aleatórias
*/