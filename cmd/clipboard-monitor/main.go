package main

import (
	"log"
	"path/filepath"
	"time"

	"github.com/andersoncorso/clipboard-monitor/internal/clipboard"
	"github.com/andersoncorso/clipboard-monitor/internal/ipchecker"
	"github.com/andersoncorso/clipboard-monitor/internal/notification"
	"github.com/andersoncorso/clipboard-monitor/internal/tray"
)

func main() {
    log.Println("Iniciando o monitoramento da área de transferência...")

    // Iniciar o system tray
    trayApp := tray.NewTrayApp()
    go trayApp.Run()

    // Carregar dados de rede
    dataFilePath := filepath.Join(".", "data.txt")
    networks, err := ipchecker.LoadNetworkData(dataFilePath)
    if err != nil {
        log.Fatalf("Erro ao carregar dados de rede: %v", err)
    }

    lastContent := ""

    for {
        content, err := clipboard.GetContent()
        if err != nil {
            log.Printf("Erro ao obter conteúdo da área de transferência: %v", err)
            time.Sleep(time.Second)
            continue
        }

        // Se o conteúdo for vazio, significa que não é texto, então ignoramos
        if content == "" {
            time.Sleep(time.Second)
            continue
        }

        if content != lastContent {
            log.Printf("Novo conteúdo de texto detectado: %s", content)

            // Verificar se o conteúdo é um IP válido e está na lista de redes
            if info := ipchecker.CheckIP(content, networks); info != "" {
                notificationMessage := "IP encontrado: " + content + "\n" + info
                err := notification.Show("Clipboard Monitor", notificationMessage)
                if err != nil {
                    log.Printf("Erro ao exibir notificação: %v", err)
                }
            } else {
                log.Println("IP não encontrado na lista ou conteúdo não é um IP válido")
            }

            lastContent = content
        }

        time.Sleep(time.Second)
    }
}