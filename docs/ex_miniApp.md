# Guia de Implementação: WebApp Telegram com API HTTP Direta

Este guia explica como implementar e executar o WebApp Telegram usando chamadas HTTP diretas à API do Telegram, sem depender de bibliotecas externas como GoBotKit.

## Estrutura de Arquivos

```
webapp-telegram/
├── main.go                 # Código principal do servidor
├── go.mod                  # Dependências Go
├── templates/              # Templates HTML
│   └── index.html          # Interface do WebApp
└── static/                 # Arquivos estáticos (opcional)
    ├── css/
    ├── js/
    └── img/
```

## Pré-requisitos

- Go 1.16 ou superior
- Token de bot do Telegram (obtido via [@BotFather](https://t.me/BotFather))
- Servidor com HTTPS para produção (pode usar ngrok para testes)
- Eu uso zero trust cloudflare

## Melhorias Implementadas

### 1. Chamadas HTTP Diretas à API do Telegram
- Removida a dependência do GoBotKit
- Implementadas funções para verificação do token e envio de mensagens usando a API HTTP do Telegram
- Melhor tratamento de erros e logs detalhados

### 2. Frontend Otimizado
- Verificação de disponibilidade do `tg.MainButton`
- Botão de envio tipo `button` em vez de `submit`
- Função `sendFormData()` separada para evitar duplicação de código
- Tratamento de erros mais detalhado

## Passos para Execução

### 1. Configurar o Bot no BotFather

1. Converse com [@BotFather](https://t.me/BotFather) no Telegram
2. Use o comando `/mybots` e selecione seu bot
3. Vá para "Bot Settings" > "Menu Button" ou "Configure Mini App"
4. Configure a URL do seu WebApp (ex: `https://seu-servidor.com/webapp`)

### 2. Instalar Dependências

```bash
# Criar o arquivo go.mod
go mod init webapp-telegram

# Instalar a lib goBotKit
go get github.com/MTplusWebSystem/GoBotKit@v0.1.8

# Atualizar dependências
go mod tidy
```

### 3. Configurar Variáveis de Ambiente

```bash
# Token do seu bot Telegram (OBRIGATÓRIO)
export BOT_TOKEN="seu_token_aqui"

# URL do servidor (opcional, padrão: http://localhost:8080)
export SERVER_URL="https://seu-servidor.com"

# Porta do servidor (opcional, padrão: 8080)
export PORT="8080"

# Modo de desenvolvimento (pula validação de assinatura para testes)
export DEV_MODE="true"  # Remova ou defina como "false" em produção
```

### 4. Executar o Servidor

```bash
# Executar o servidor
go run main.go
```

### 5. Testar com KitDomain (Desenvolvimento)

Para testar localmente, você pode usar o zero trust cloudflare para criar um túnel HTTPS:

```bash
# Criar conta na cloudflare

opção "zero trust" > rede

opção "Túneis" > Criar um túnel > Nomeie seu túnel > Escolha seu ambiente 

opção "debian" > Instale e execute um conector >

```

Use a URL HTTPS fornecida pelo zero trust para configurar seu WebApp no BotFather.

## Características do Código

### Backend (MicroServe.go)

- **Chamadas HTTP Diretas**: Usa a biblioteca padrão `net/http` para comunicação com a API do Telegram
- **Validação de Dados**: Middleware que verifica a autenticidade dos dados enviados pelo Telegram
- **API REST**: Endpoints para receber dados do WebApp e enviar mensagens via bot
- **CORS Configurado**: Permite requisições de qualquer origem (configurável)
- **Rotas Robustas**: Inclui redirecionamentos e tratamento de erros
- **Modo de Desenvolvimento**: Facilita testes locais sem validação de assinatura

### Frontend (index.html)

- **Endpoint Dinâmico**: Adapta-se automaticamente ao domínio do servidor
- **Tema Adaptativo**: Usa as cores do tema do Telegram do usuário
- **Responsivo**: Funciona em qualquer tamanho de tela
- **Botão Principal**: Utiliza o tg.MainButton do Telegram para envio
- **Tratamento de Erros**: Exibe mensagens de erro e sucesso
- **Dados do Usuário**: Exibe nome e foto do usuário automaticamente

## Nota Importante sobre a API do Telegram

A API do Telegram WebApp está em constante evolução. Este código utiliza `tg.MainButton` em vez de `tg.BottomButton`, que foi descontinuado em versões recentes. Se você encontrar erros relacionados a componentes indefinidos, verifique a documentação mais recente do Telegram para possíveis mudanças na API.

## Fluxo de Funcionamento

1. O usuário interage com o bot e clica no botão que abre o WebApp
2. O Telegram abre o WebApp dentro do aplicativo
3. O WebApp se inicializa e obtém dados do usuário
4. O usuário preenche o formulário e envia
5. O servidor valida os dados recebidos
6. O servidor envia uma mensagem diretamente via API HTTP do Telegram


## Personalização

Para adaptar este exemplo ao seu caso de uso:

1. **Modifique o Frontend**: Altere o HTML/CSS/JS em `templates/index.html` para sua interface
2. **Adicione Funcionalidades**: Implemente novos endpoints na API conforme necessário
3. **Integre com Banco de Dados**: Adicione persistência de dados se necessário
4. **Expanda a Integração com o Bot**: Adicione mais comandos e interações no bot

## Considerações para Produção

1. **Desative o Modo de Desenvolvimento**: Remova ou defina `DEV_MODE="false"` em produção
2. **Use HTTPS**: WebApps do Telegram só funcionam com HTTPS em produção
3. **Configure um Domínio Fixo**: Para produção, é recomendável ter um domínio fixo
4. **Implemente Logging**: Adicione logging mais detalhado para depuração
5. **Considere Escalabilidade**: Para aplicações maiores, considere separar o backend em microserviços

## Solução de Problemas Comuns

### Erro "Erro ao enviar mensagem"

Este erro pode ocorrer por vários motivos:

1. **Token do bot inválido**: Verifique se o token está correto e atualizado
2. **Bot não iniciado**: O usuário precisa iniciar uma conversa com o bot antes de receber mensagens
3. **Problemas de rede**: Verifique se o servidor tem acesso à API do Telegram
4. **Permissões do bot**: Certifique-se de que o bot não foi bloqueado pelo usuário

### Botão de Enviar Não Funciona

Se o botão de enviar não estiver funcionando:

1. **Verifique o console do navegador** para erros JavaScript
2. **Confirme que o Telegram WebApp está carregando corretamente**
3. **Teste em diferentes dispositivos** (alguns recursos podem variar entre plataformas)

## Verificação do Token do Bot

O servidor agora verifica automaticamente se o token do bot é válido na inicialização usando uma chamada direta à API do Telegram. Se o token for inválido, você verá uma mensagem de aviso nos logs, mas o servidor continuará funcionando para permitir testes da interface.

Para verificar manualmente se o token do bot é válido, você pode usar:

```bash
curl -X GET "https://api.telegram.org/bot$BOT_TOKEN/getMe"
```

Se o token for válido, você receberá informações sobre o bot. Caso contrário, receberá um erro.

## Recursos Adicionais

- [Documentação Oficial de Mini Apps do Telegram](https://core.telegram.org/bots/webapps)
- [Documentação da API do Telegram](https://core.telegram.org/bots/api)
- [Uso rápido](https://github.com/MTplusWebSystem/GoBotKit/cmd/test_miniApp.go)
- [Código do MicroServer](https://github.com/MTplusWebSystem/GoBotKit/botkit/MicroServer.go)
