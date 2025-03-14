package tray

import (
	"fmt"
	"os"
	"path/filepath"

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
    iconData, err := getIcon("icon.ico")
    if err != nil {
        fmt.Printf("Erro ao carregar o ícone: %v\n", err)
        // Use um ícone padrão ou continue sem um ícone
        // systray.SetIcon(defaultIcon)
    } else {
        systray.SetIcon(iconData)
    }

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
// Tenta múltiplos locais possíveis e retorna um erro se não encontrar o ícone.
func getIcon(iconName string) ([]byte, error) {
    possiblePaths := []string{
        iconName,
        filepath.Join("assets", iconName),
        filepath.Join("build", "assets", iconName),
    }

    // Adiciona o diretório do executável à lista de caminhos possíveis
    if exePath, err := os.Executable(); err == nil {
        exeDir := filepath.Dir(exePath)
        possiblePaths = append(possiblePaths,
            filepath.Join(exeDir, iconName),
            filepath.Join(exeDir, "assets", iconName),
            filepath.Join(exeDir, "build", "assets", iconName),
        )
    }

    for _, path := range possiblePaths {
        if b, err := os.ReadFile(path); err == nil {
            return b, nil
        }
    }

    return nil, fmt.Errorf("não foi possível encontrar o ícone: %s", iconName)
}