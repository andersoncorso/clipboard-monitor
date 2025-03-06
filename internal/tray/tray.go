package tray

import (
	"os"

	"github.com/getlantern/systray"
)

type TrayApp struct {
    quitChan chan struct{}
}

func NewTrayApp() *TrayApp {
    return &TrayApp{
        quitChan: make(chan struct{}),
    }
}

func (app *TrayApp) Run() {
    systray.Run(app.onReady, app.onExit)
}

func (app *TrayApp) Quit() {
    systray.Quit()
    <-app.quitChan
}

func (app *TrayApp) onReady() {
    systray.SetIcon(getIcon("icon.ico"))
    systray.SetTitle("Clipboard Monitor")
    systray.SetTooltip("Monitorando a área de transferência")

    mQuit := systray.AddMenuItem("Sair", "Sair do aplicativo")

    go func() {
        <-mQuit.ClickedCh
        systray.Quit()
    }()
}

func (app *TrayApp) onExit() {
    close(app.quitChan)
}

func getIcon(s string) []byte {
    b, err := os.ReadFile(s)
    if err != nil {
        panic(err)
    }
    return b
}