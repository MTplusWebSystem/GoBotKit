package botkitv2

import "fmt"

type MiniApp struct {
	Config *ConfigGlobal
}

func NewMiniApp(cfg *ConfigGlobal) *MiniApp {
	return &MiniApp{
		Config: cfg,
	}
}

func (miniApp *MiniApp) StartMiniApp() {
	fmt.Println("Seu MiniApp",miniApp.Config.Token)
}