package botkit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/MTplusWebSystem/GoBotKit/requests"
	"github.com/MTplusWebSystem/GoBotKit/utils"
)


func (b *BotInit) ReplyToPhotoButton( photoPath string, replyMarkup interface{}) error {
	text := ""
    apiUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", b.Token)

    if photoPath != "" {
        photoFile, err := os.Open(photoPath)
        utils.NilError("",err)

        defer photoFile.Close()

        body := &bytes.Buffer{}
        writer := multipart.NewWriter(body)

        photoPart, err := writer.CreateFormFile("photo", filepath.Base(photoPath))
        utils.NilError("",err)
        
        _, err = io.Copy(photoPart, photoFile)
        utils.NilError("",err)
        
        params := map[string]string{
            "chat_id": strconv.Itoa(b.ChatID),
            "text":    text,
        }

        for key, value := range params {
            _ = writer.WriteField(key, value)
        }


        replyMarkupJSON, err := json.Marshal(replyMarkup)
        utils.NilError("",err)

        _ = writer.WriteField("reply_markup", string(replyMarkupJSON))

        err = writer.Close()
        utils.NilError("",err)
		contenttype := writer.FormDataContentType()
        requests.POSTMultipart(apiUrl,contenttype, body)
    } else {
        params := url.Values{}
        params.Set("chat_id", strconv.Itoa(b.ChatID))
        params.Set("text", text)

        replyMarkupJSON, err := json.Marshal(replyMarkup)
        utils.NilError("",err)

        params.Set("reply_markup", string(replyMarkupJSON))

        resp, err := http.PostForm(apiUrl, params)
        utils.NilError("",err) 

        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
            return fmt.Errorf("Erro ao enviar a resposta. Código de status: %d", resp.StatusCode)
        }
    }

    return nil
}
