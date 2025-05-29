package main

import (
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/andersoncorso/clipboard-monitor/internal/clipboard"
	"github.com/andersoncorso/clipboard-monitor/internal/ipchecker"
	"github.com/andersoncorso/clipboard-monitor/internal/notification"
	"github.com/andersoncorso/clipboard-monitor/internal/tray"
	"github.com/gofrs/flock"
)

func main() {
    // Cria um arquivo de bloqueio para garantir instância única
    lockFile := filepath.Join(os.TempDir(), "clipboard-monitor.lock")
    fileLock := flock.New(lockFile)

    // Tenta obter o bloqueio
    locked, err := fileLock.TryLock()
    if err != nil {
        log.Fatalf("Erro ao tentar obter o bloqueio: %v", err)
    }

    // Se não conseguir obter o bloqueio, outra instância já está em execução
    if !locked {
        log.Println("Outra instância do aplicativo já está em execução.")
        notification.Show("Clipboard Monitor", "O aplicativo já está em execução.")
        return
    }

    // Libera o bloqueio quando o programa terminar
    defer fileLock.Unlock()

    log.Println("Iniciando o monitoramento da área de transferência...")

    // Inicializa e executa o ícone na área de notificação (system tray)
    // Isso é executado em uma goroutine separada para não bloquear o programa principal
    trayApp := tray.NewTrayApp()
    go trayApp.Run()

    // Carrega os dados de rede do arquivo data.txt
    // Esse arquivo contém as informações de IP e suas descrições
    dataFilePath := filepath.Join("build", "data", "data.txt")
    networks, err := ipchecker.LoadNetworkData(dataFilePath)
    if err != nil {
        // Se houver um erro ao carregar os dados, o programa é encerrado
        log.Fatalf("Erro ao carregar dados do arquivo: %v", err)
    }

    lastContent := ""

    for {
        // Obtém o conteúdo atual da área de transferência
        content, err := clipboard.GetContent()
        if err != nil {
            log.Printf("Erro ao obter conteúdo da área de transferência: %v", err)
            time.Sleep(time.Second)
            continue
        }

        content = strings.TrimSpace(content)

        // Adiciona logs detalhados para depuração
        log.Printf("Último conteúdo: %q", lastContent)
        log.Printf("Conteúdo atual: %q", content)

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
            } else if isHostname(content) {
                // Se for um hostname, resolve o IP
                ips, err := net.LookupIP(content)
                if err != nil {
                    log.Printf("Erro ao resolver hostname: %v", err)
                    continue
                }
                for _, ip := range ips {
                    if info := ipchecker.CheckIP(ip.String(), networks); info != "" {
                        // Se o IP do hostname estiver na lista, prepara e exibe uma notificação
                        notificationMessage := "Hostname resolvido: " + content + "\nIP encontrado: " + ip.String() + "\n" + info
                        err := notification.Show("Clipboard Monitor", notificationMessage)
                        if err != nil {
                            log.Printf("Erro ao exibir notificação: %v", err)
                        }
                        break
                    }
                }
            }

            // Atualiza o último conteúdo processado
            lastContent = content
        }

        // Pausa por um segundo antes de verificar novamente
        // Isso evita uso excessivo de CPU
        time.Sleep(time.Second)
    }
}

// Verifica se o conteúdo é um hostname válido
func isHostname(content string) bool {
    return strings.HasPrefix(content, "LP") || strings.HasPrefix(content, "DT") || strings.HasPrefix(content, "FILA")
}
