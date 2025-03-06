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
    // Inicia o log do programa
    log.Println("Iniciando o monitoramento da área de transferência...")

    // Inicializa e executa o ícone na área de notificação (system tray)
    // Isso é executado em uma goroutine separada para não bloquear o programa principal
    trayApp := tray.NewTrayApp()
    go trayApp.Run()

    // Carrega os dados de rede do arquivo data.txt
    // Esse arquivo contém as informações de IP e suas descrições
    dataFilePath := filepath.Join(".", "data.txt")
    networks, err := ipchecker.LoadNetworkData(dataFilePath)
    if err != nil {
        // Se houver um erro ao carregar os dados, o programa é encerrado
        log.Fatalf("Erro ao carregar dados do arquivo: %v", err)
    }

    // Variável para armazenar o último conteúdo da área de transferência
    // Isso é usado para evitar processamento repetido do mesmo conteúdo
    lastContent := ""

    // Loop principal do programa
    for {
        // Obtém o conteúdo atual da área de transferência
        content, err := clipboard.GetContent()
        if err != nil {
            // Se houver um erro ao obter o conteúdo, registra o erro e continua o loop
            log.Printf("Erro ao obter conteúdo da área de transferência: %v", err)
            time.Sleep(time.Second)
            continue
        }

        // Se o conteúdo for vazio, significa que não é texto, então é ignorado
        if content == "" {
            time.Sleep(time.Second)
            continue
        }

        // Verifica se o conteúdo é diferente do último processado
        if content != lastContent {
            log.Printf("Novo conteúdo de texto detectado: %s", content)

            // Verifica se o conteúdo é um IP válido e está na lista de redes
            if info := ipchecker.CheckIP(content, networks); info != "" {
                // Se for um IP válido, prepara e exibe uma notificação
                notificationMessage := "IP encontrado: " + content + "\n" + info
                err := notification.Show("Clipboard Monitor", notificationMessage)
                if err != nil {
                    log.Printf("Erro ao exibir notificação: %v", err)
                }
            } else {
                // Se não for um IP válido ou não estiver na lista, registra isso no log
                log.Println("IP não encontrado na lista ou conteúdo não é um IP válido")
            }

            // Atualiza o último conteúdo processado
            lastContent = content
        }

        // Pausa por um segundo antes de verificar novamente
        // Isso evita uso excessivo de CPU
        time.Sleep(time.Second)
    }
}