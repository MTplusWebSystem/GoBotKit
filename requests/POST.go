package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/MTplusWebSystem/GoBotKit/utils"
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
	utils.NilError("",err)

	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(jsonValue))
	utils.NilError("",err)
	if ContentType != "" {
		req.Header.Set("Content-Type", ContentType)
	}else{
		req.Header.Set("Content-Type", "application/json")
	}
	
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status: " + resp.Status)
	}
	utils.NilError("",err)

	return resp, nil
}

func POSTMultipart(url, ContentType string, body *bytes.Buffer ) error {
	writer := multipart.NewWriter(body)

	writer.Close()

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return fmt.Errorf("erro ao criar requisição HTTP: %v", err)
	}
	req.Header.Set("Content-Type", ContentType )

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erro ao enviar requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erro ao enviar a foto. Código de status: %d", resp.StatusCode)
	}

	return nil
}