
# GoBotKit 🤖

Biblioteca simples e poderosa para criar bots no Telegram com Go.

## 📦 Instalação

```bash
go get https://github.com/MTplusWebSystem/GoBotKit/pkg
```

## 🚀 Funcionalidades Principais

- 🔹 Envio de mensagens
- 🔹 Criação de botões interativos
- 🔹 Manipulação de documentos, arquivos e mídias
- 🔹 Tratamento de eventos (mensagens, cliques e comandos)
- 🔹 Obtenção de arquivos enviados para o bot

## 🧠 Estrutura dos Objetos

### Documento (Document)

```go
type Document struct {
    Status        bool
    FileName      string
    MimeType      string
    FileID        string
    FileUniqueID  string
    FileSize      int
}
```

## 🔧 Exemplos de Uso

### Enviar Mensagem Simples

```go
bot.SendMessages("Olá, mundo!")
```

### Enviar Mensagem com Markdown

```go
bot.SendMessages("```go\nfmt.Println(\"Hello World\")```", "markdown")
```

### Enviar Mensagem com Botões

```go
layout := [][]string{
    {"Botão 1", "Botão 2"},
    {"Botão 3"},
}
bot.SendButton("Escolha uma opção:", layout)
```

### Obter Arquivo Enviado

```go
switch bot.Document.FileName {
case "exemplo.pdf":
    bot.GetFile("./storage") // Salvar na pasta storage
}
```

> ⚠️ **Atenção:** `GetFile()` não deve ser usado dentro de `bot.Handler()`. Utilize após o recebimento do arquivo.

## 🔥 Manipulação de Eventos

```go
bot.Handler(func() {
    if bot.Text == "/start" {
        bot.SendMessages("Bem-vindo!")
    }
})
```

## 🗂 Estrutura do Projeto

```
botkit/
 ├── DeleteMessage.go
 ├── EventHandler.go
 ├── SendMessages.go
 ├── SendButton.go
 ├── ...
requests/
 ├── GET.go
 ├── POST.go
utils/
 ├── system.go
```

## 📜 Licença

MIT License
