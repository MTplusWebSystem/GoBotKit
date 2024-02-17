package requests

import (
	"net/http"

	"github.com/MTplusWebSystem/GoBotKit/system"
)

/*
Pacote requests inicilmente vai estar disponivel GET, POST

funções para requisição GET

GET : para requisição GET
ReadBody : para ler o corpo da requisição
*/

func GET(url string) (*http.Response, error){
	req , err := http.Get(url)
	system.NilError("Erro na solicitação",err)
	return req, nil
}