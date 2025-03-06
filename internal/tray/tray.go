package tray

import (
	"os"

	"github.com/getlantern/systray"
)

// TrayApp representa a aplicação na área de notificação.
type TrayApp struct {
    quitChan chan struct{} // Canal para sinalizar quando a aplicação deve ser encerrada
}

// NewTrayApp cria e retorna uma nova instância de TrayApp.
func NewTrayApp() *TrayApp {
    return &TrayApp{
        quitChan: make(chan struct{}),
    }
}

// Run inicia a aplicação na área de notificação.
// Esta função bloqueia até que a aplicação seja encerrada.
func (app *TrayApp) Run() {
    systray.Run(app.onReady, app.onExit)
}

// Quit encerra a aplicação na área de notificação.
// Esta função bloqueia até que a aplicação seja completamente encerrada.
func (app *TrayApp) Quit() {
    systray.Quit()
    <-app.quitChan // Espera até que onExit() sinalize que a aplicação foi encerrada
}

// onReady é chamada quando a área de notificação está pronta.
// Configura o ícone, título, tooltip e menu da aplicação.
func (app *TrayApp) onReady() {
    systray.SetIcon(getIcon("icon.ico"))
    systray.SetTitle("Clipboard Monitor")
    systray.SetTooltip("Monitorando a área de transferência")

    mQuit := systray.AddMenuItem("Sair", "Sair do aplicativo")

    // Goroutine para lidar com o clique no item de menu "Sair"
    go func() {
        <-mQuit.ClickedCh
        systray.Quit()
    }()
}

// onExit é chamada quando a aplicação na área de notificação está sendo encerrada.
func (app *TrayApp) onExit() {
    close(app.quitChan) // Sinaliza que a aplicação foi encerrada
}

// getIcon lê um arquivo de ícone e retorna seus bytes.
// Panics se não conseguir ler o arquivo (isso pode ser melhorado para retornar um erro).
func getIcon(s string) []byte {
    b, err := os.ReadFile(s)
    if err != nil {
        panic(err)
    }
    return b
}