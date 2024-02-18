package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"github.com/MTplusWebSystem/GoBotKit/system"
)

/*
Pacote requests inicilmente vai estar disponivel GET, POST

funções para requisição POST

POST: para requisição POST

POSTmultipart : para requisição multipart
*/

func POST(Url ,ContentType string, params interface{}) (*http.Response, error) {
	client := &http.Client{}

	jsonValue, err := json.Marshal(params)
	system.NilError("",err)

	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(jsonValue))
	system.NilError("",err)
	if ContentType != "" {
		req.Header.Set("Content-Type", ContentType)
	}else{
		req.Header.Set("Content-Type", "application/json")
	}
	
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status: " + resp.Status)
	}
	system.NilError("",err)

	return resp, nil
}

