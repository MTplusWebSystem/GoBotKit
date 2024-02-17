package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
)

func NilError(msg string , err error) {
	if err != nil {
		fmt.Println(msg,err)
		return
	}
}

func Scan(path string) []byte {
	content , err := ioutil.ReadFile(path)
	NilError(fmt.Sprint("Erro ao ler o arquivo %s", path), err)
	return content
}

func Regex(params, text string) string {
	re := regexp.MustCompile(params)
    match := re.FindStringSubmatch(text)
    if len(match) > 1 {
        return match[1]
    }
    return ""
}

func WriteJSON(filename string, data interface{}) error {
	jsonData, err := json.Marshal(data)
    NilError("Erro ao ler sua interface", err)

    err = ioutil.WriteFile(filename, jsonData, 0644)
	NilError(fmt.Sprintf("Erro ao escrever no path:%s",filename), err)

    return nil
}
/*
funções para serem criadas

NilError: simplificar as mensagens de erros e diminuir linhas de código |ok

Scan: para leitura simples de arquivos |ok

Regex: Uma forma simples de aplicar regex |ok

WriteJSON: para escrita de arquivos JSON |ok

Random: para geração de valores aleatórios

KeyGenerator: para geração de chaves aleatórias
*/