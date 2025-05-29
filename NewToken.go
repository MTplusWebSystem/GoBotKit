package botkitv2

type ConfigGlobal struct {
	Token string
	ChatID int
}

func NewToken(token string) *ConfigGlobal {
	return &ConfigGlobal{
		Token:token,
	}
}