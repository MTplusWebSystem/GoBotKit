package botkit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/MTplusWebSystem/GoBotKit/requests"
	"github.com/MTplusWebSystem/GoBotKit/system"
)



func (b *BotInit) GetFile(filePATH ...string) {

		url := fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", b.Token, b.Document.FileID)

		response, err := requests.GET(url)
		system.NilError("", err)
		defer response.Body.Close()
	
		var fileData struct {
			OK     bool `json:"ok"`
			Result struct {
				FilePath string `json:"file_path"`
			} `json:"result"`
		}
		err = json.NewDecoder(response.Body).Decode(&fileData)
		system.NilError("Erro ao decodificar resposta JSON:", err)
	
	
		fileDownloadURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", b.Token, fileData.Result.FilePath)
	
		responseFile, err := requests.GET(fileDownloadURL)
		system.NilError("Erro ao baixar o arquivo:", err)
		defer responseFile.Body.Close()
	
		fileContents, err := ioutil.ReadAll(responseFile.Body)
		system.NilError("Erro ao ler conteúdo do arquivo:", err)
		save := "./"
		if len(filePATH) > 0{
			save = filePATH[0]+"/"+b.Document.FileName
		}
		err = ioutil.WriteFile(save, fileContents, 0644)
		if err != nil {
			fmt.Println("Erro ao salvar o arquivo:", err)
			return
		}
	
	}
	

