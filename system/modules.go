package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func Random(params string, custom bool) string{
	rand.Seed(time.Now().UnixNano())

	letrasMin := "abcdefghijklmnopqrstuvwxyz"
	letrasMai := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letrasMinMai := letrasMin + letrasMai
	if custom {
		if params[0:1] != "a" && params[0:1] != "A" {
			op := ""
			parts := strings.Split(params, "-")
			if len(parts) == 2 {
				start, _ := strconv.Atoi(parts[0])
				final, _ := strconv.Atoi(parts[1])
				value :=  rand.Intn(final-start+1) + start
				op = strconv.Itoa(value)
			}
			return op
		}
	}
	switch params{
		case "a-z":
            indice := rand.Intn(len(letrasMin))
			generated := string(letrasMin[indice])
            return generated
        case "A-Z":
            indice := rand.Intn(len(letrasMai))
			generated := string(letrasMai[indice])
            return generated
        case "a-Z":
            indice := rand.Intn(len(letrasMinMai))
			generated := string(letrasMinMai[indice])
            return generated
        case "0-9":
            indice := rand.Intn(10)
            return strconv.Itoa(indice)
        default:
            return "INVALID"
	}
}

func KeyGenerator(token int) string {
	
	key := make([]string, 0)

	for i := 0; i < token; i++ {
        par := ""
        for j := 0; j < token; j++ {
            if j <= 1 || j >= 3 {
                par += Random("a-Z", false)
            } else {
                par += Random("0-9", false)
            }
        }
        key = append(key, par)
    }
	form := fmt.Sprintf("%s-%s-%s-%s",key[0],key[1],key[2],key[3])
	return form
}
/*
funções para serem criadas

NilError: simplificar as mensagens de erros e diminuir linhas de código |ok

Scan: para leitura simples de arquivos |ok

Regex: Uma forma simples de aplicar regex |ok

WriteJSON: para escrita de arquivos JSON |ok

Random: para geração de valores aleatórios |ok

KeyGenerator: para geração de chaves aleatórias
*/