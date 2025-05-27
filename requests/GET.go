package requests

import (
	"io/ioutil"
	"net/http"
	"github.com/MTplusWebSystem/GoBotKit/utils"
)

/*
Pacote requests inicilmente vai estar disponivel GET, POST

funções para requisição GET

GET : para requisição GET
ReadBody : para ler o corpo da requisição
*/

func GET(url string) (*http.Response, error){
	req , err := http.Get(url)
	utils.NilError("Erro na solicitação",err)
	return req, nil
}

func ReadBody(response *http.Response) []byte {
	body, err := ioutil.ReadAll(response.Body)
	utils.NilError("Erro ao ler o corpo da resposta:", err)
	defer response.Body.Close()
	return body
}