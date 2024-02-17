
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

