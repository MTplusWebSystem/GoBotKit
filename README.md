# Pacote Botkit :robot:

### Funções atuais
- **SendMessages**:
```go
go func() {
	bot.Handler("callback_query",func(event string) {
		fmt.Println(bot.ChatID)
		fmt.Println("tipo:callback_query(", event,")")
		if event == "!cadastro" {
			bot.SendMessages("Nome: \nSobre nome: \n idade: \n")
		}
     })
}()

```
- **SendButton**:
```go
go func() {
	bot.Handler("commands",func(event string) {
		fmt.Println("tipo:commands(",event,")")
		if event == "/menu" {
			layout := map[string]interface{}{
				"inline_keyboard": [][]map[string]interface{}{
					{
						{"text": "Cadastrar", "callback_data": "!cadastro"},
						{"text": "Sair", "callback_data": "!sair"},
					},
				},
			}
		bot.SendButton("Bem-vindo ao menu",layout)
		}
	})
}()
```
- **ReplyToMessage**:
```go
go func() {
	bot.Handler("callback_query",func(event string) {
		fmt.Println(bot.ChatID)
		fmt.Println("tipo:callback_query(", event,")")
		if event == "!cadastro" {
			bot.SendMessages("Nome: \nSobre nome: \n idade: \n")
		}
		if event == "!sair"{
			bot.ReplyToMessage(bot.QueryMessageID, "Tem certeza ?\n")
			layout := map[string]interface{}{
				"inline_keyboard": [][]map[string]interface{}{
					{
						{"text": "Cancelar", "callback_data": "!cancelar"},
						{"text": "Continuar", "callback_data": "!continuar"},
					},
				},
			}
			bot.SendButton("ainda não terminou o cadastro",layout)
		}
	})
}()
```

### nova função

- **SendPhoto**:

```go
go func() {
	bot.Handler("commands",func(event string) {
		fmt.Println("tipo:commands(",event,")")
		if event == "/menu" {
			layout := map[string]interface{}{
				"inline_keyboard": [][]map[string]interface{}{
					{
						{"text": "Cadastrar", "callback_data": "!cadastro"},
						{"text": "Sair", "callback_data": "!sair"},
					},
				},
			}
			bot.SendButton("Bem-vindo ao menu",layout)
		} else if event == "/start" {
			bot.SendPhoto("./boas-vinda.jpg","Olá sejá bem-vindo")
		}
	})
}()
```

# Pacote Botkit :robot:

Este pacote oferece funções para criação e interação com bots do Telegram.

## Estrutura `BotInit` :gear:

Esta estrutura é usada para inicializar e interagir com um bot do Telegram.

### Campos:

- `Token`: Token de autenticação do bot.
- `UpdateID`: ID da última atualização recebida.
- `MessageID`: ID da última mensagem recebida.
- `ChatID`: ID do chat atual.
- `CallbackID`: ID do callback atual.
- `Username`: Nome de usuário do bot.
- `Text`: Texto da mensagem recebida.
- `CallbackQuery`: Query do callback.
- `QueryMessageID`: ID da mensagem associada à query.
- `ID`: ID do usuário ou do chat.

## Funções Disponíveis:

- `GetUpdates`: Obtém as atualizações mais recentes do bot.

## Como Usar :rocket:

1. Importe o pacote no seu código Go:

   ```go
   import "github.com/MTplusWebSystem/GoBotKit/botkit"
   ```

2. Utilize a estrutura `BotInit` para inicializar o bot e interagir com ele.

## Exemplos de Uso :bulb:

```go
package main

import (
	"fmt"

	"github.com/MTplusWebSystem/GoBotKit/botkit"
)


func main() {
	bot := botkit.BotInit{
		Token: "seu-token-aqui",
	}
	for {
		if bot.ReceiveData(){
			go func() {
				bot.Handler("callback_query",func(event string) {
					fmt.Println("tipo:callback_query(", event,")")
				})
			}()
			go func() {
				bot.Handler("commands",func(event string) {
					fmt.Println("tipo:commands(",event,")")
				})
			}()

			go func() {
				bot.Handler("messages", func(event string) {
					fmt.Println("tipo:messages(",event,")")
					if event == "olá"{
						bot.SendMessages("Olá tudo-bem")
					}
				})
			}()
		}
	}
}

```



---Para utilizar as funcionalidades do pacote, você pode implementar um loop para receber as atualizações do bot e processá-las conforme necessário. Por exemplo:

```go
// Exemplo de loop para processar atualizações do bot
for {
    updates, err := bot.GetUpdates()
    if err != nil {
        fmt.Println("Erro ao obter atualizações:", err)
        return
    }

    for _, update := range updates {
        // Implemente aqui a lógica para processar cada atualização
    }
}
```

Certifique-se de revisar a documentação para garantir que as instruções de uso estejam atualizadas com a implementação mais recente do pacote.


# Pacote Requests :arrow_up:

Este pacote fornece funções para realizar requisições HTTP do tipo GET.

## Funções Disponíveis

### `GET` :rocket:

Realiza uma requisição HTTP GET para o URL fornecido.

### `ReadBody` :book:

Lê o corpo da resposta de uma requisição HTTP GET.

### `POST` :rocket:

Realiza uma requisição HTTP POST para o URL fornecido.

### `POSTmultipart` :rocket:

Realiza uma requisição HTTP POST multipart para o URL fornecido.

## Como Usar :rocket:


1. Importe o pacote no seu código Go:

```go
   import "github.com/seu-usuario/seu-repositorio/requests"
```

# Exemplos de Uso :bulb:

```go
// Exemplo de utilização da função GET
response, err := requests.GET("https://exemplo.com/api/dados")
if err != nil {
    fmt.Println("Erro na requisição GET:", err)
    return
}

// Exemplo de utilização da função ReadBody
body := requests.ReadBody(response) 
fmt.Println("Corpo da resposta:", string(body))

// Exemplo de utilização da função POST
params := map[string]interface{}{"nome": "exemplo", "idade": 30}
response, err := requests.POST("https://exemplo.com/api/dados", "application/json", params)
if err != nil {
    fmt.Println("Erro na requisição POST:", err)
    return
}
```



# Pacote System :v0.1

Este pacote fornece uma coleção de funções úteis para diversas operações no sistema.

## Funções Disponíveis

### `NilError` :warning:

Função para simplificar mensagens de erros e reduzir linhas de código.

### `Scan` :file_folder:

Realiza a leitura simples de arquivos.

### `Regex` :mag_right:

Fornece uma maneira simples de aplicar regex.

### `WriteJSON` :page_with_curl:

Permite a escrita de arquivos JSON.

### `Random` :game_die:

Gera valores aleatórios de acordo com os parâmetros fornecidos.

### `KeyGenerator` :key:

Gera chaves aleatórias para uso em sistemas de autenticação.

## Como Usar :rocket:

1. Importe o pacote no seu código Go:

 ```go
   import "github.com/MTplusWebSystem/GoBotKit/system"
 ```

### Exemplo de uso


```go
package main

import (
	"fmt"
	"github.com/MTplusWebSystem/GoBotKit/system"
)

func main() {

//Scan , tenha certeza de criar o arquivo antes
	data := string(system.Scan("./arquivo.txt"))
	fmt.Println(data)
	
//Regex, a expresão regular deve ser a primeira
	resultado := system.Regex("[0-9]+", "texto123")
	fmt.Println(resultado)

//WriteJSON, tenha certeza da relação entre o código onde está sendo executado e onde vai salvar
	dados := map[string]interface{}{"nome": "Exemplo", "idade": 30}
	err := system.WriteJSON("dados.json", dados)
	system.NilError("Erro ao escrever arquivo JSON:", err)
	
//Ramdom, custom inicialmente é para valores maior que 9 exemplo: 10-300
	aleatorio := system.Random("0-9", false)
	fmt.Println(aleatorio)

//KeyGererator, está configurado para gerar key em divisivos de 2 exemplo: 2,4,6,8...
	chave := system.KeyGenerator(4)
	fmt.Println(chave)
}
```

