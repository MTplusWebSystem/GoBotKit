package botkit

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"github.com/MTplusWebSystem/GoBotKit/requests"
	"github.com/MTplusWebSystem/GoBotKit/utils"
)


func (b *BotInit) SendPhoto(photoPath, caption string) error {
	photoFile, err := os.Open(photoPath)
	utils.NilError("",err)
	
	defer photoFile.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("photo", photoPath)
	utils.NilError("",err)

	_, err = io.Copy(part, photoFile)
	utils.NilError("",err)

	err = writer.WriteField("caption", caption)
	utils.NilError("",err)

	err = writer.Close()
	utils.NilError("",err)

	Url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto?chat_id=%d", b.Token, b.ChatID)

	var contenttype string = writer.FormDataContentType()
	requests.POSTMultipart(Url,contenttype,body)
	
	return nil
}