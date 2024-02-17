package system

import (
	"fmt"
	"io/ioutil"
)

func NilError(msg string , err error) {
	if err != nil {
		fmt.Println(msg,err)
	}
}

func Scan(path string) []byte {
	content , err := ioutil.ReadFile(path)
	NilError(fmt.Sprint("Erro ao ler o arquivo %s", path), err)
	return content
}

/*
funções para serem criadas

NilError: simplificar as mensagens de erros e diminuir linhas de código |ok

Scan: para leitura simples de arquivos |ok

Regex: Uma forma simples de aplicar regex

WriteJSON: para escrita de arquivos JSON

Random: para geração de valores aleatórios

KeyGenerator: para geração de chaves aleatórias
*/