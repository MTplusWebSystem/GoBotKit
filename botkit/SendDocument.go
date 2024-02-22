package botkit

import (
    "bytes"
    "fmt"
    "io"
    "mime/multipart"
    "os"
    "github.com/MTplusWebSystem/GoBotKit/requests"
    "github.com/MTplusWebSystem/GoBotKit/system"
)

var nameTypes = map[string]string{
    "pdf":    "application/pdf",
    "zip":    "application/zip",
    "tar.gz": "application/x-tar-gz",
    "sql":    "application/sql",
    "txt":    "text/plain",
    "tar":    "application/x-tar",
    "json":   "application/json",
}

func (b *BotInit) SendDocument(filePath, caption, nametype string) error {
    file, err := os.Open(filePath)
    system.NilError("", err)
    defer file.Close()

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    part, err := writer.CreateFormFile("document", filePath)
    system.NilError("", err)

    _, err = io.Copy(part, file)
    system.NilError("", err)

    err = writer.WriteField("caption", caption)
    system.NilError("", err)
   
    err = writer.WriteField("type", nameTypes[nametype]) 

    err = writer.Close()
    system.NilError("", err)

    Url := fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument?chat_id=%d", b.Token, b.ChatID)

    contentType := writer.FormDataContentType()
    requests.POSTMultipart(Url, contentType, body)

    return nil
}
