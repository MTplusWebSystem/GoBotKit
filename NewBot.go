package botkitv2

import "fmt"

type Bot struct {
	Config *ConfigGlobal
}

// NewBot cria uma nova instância do Bot
func NewBot(cfg *ConfigGlobal) *Bot {
	fmt.Println("Criando conexão..., token:", cfg.Token)
	return &Bot{
		Config: cfg,
	}
}

func (b *Bot) SendMsg(msg string)  {
	fmt.Println("mensagem -> %s enviada -> %d", msg, b.Config.ChatID )
}