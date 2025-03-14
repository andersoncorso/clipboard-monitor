package ipchecker

import (
	"bufio"
	"net"
	"os"
	"strings"
)

// NetworkInfo representa uma rede e sua informação associada.
type NetworkInfo struct {
    Network string // Endereço de rede em formato CIDR (ex: "192.168.0.0/24")
    Info    string // Informação associada à rede
}

// LoadNetworkData carrega os dados de rede a partir de um arquivo.
// O arquivo deve conter linhas no formato "rede | informação".
// Retorna uma slice de NetworkInfo e um erro, se houver.
func LoadNetworkData(filename string) ([]NetworkInfo, error) {
    // Abre o arquivo
    file, err := os.Open(filename)
    if err != nil {
        return nil, err // Retorna o erro se não puder abrir o arquivo
    }
    defer file.Close() // Garante que o arquivo será fechado ao final da função

    var networks []NetworkInfo
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        // Divide a linha em duas partes usando "|" como separador
        parts := strings.SplitN(line, "|", 2)
        if len(parts) == 2 {
            // Adiciona uma nova NetworkInfo à slice, removendo espaços em branco
            networks = append(networks, NetworkInfo{
                Network: strings.TrimSpace(parts[0]),
                Info:    strings.TrimSpace(parts[1]),
            })
        }
    }

    // Verifica se houve erro durante a leitura do arquivo
    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return networks, nil
}

// CheckIP verifica se um IP está dentro de alguma das redes fornecidas.
// Retorna a informação associada à rede que contém o IP, ou uma string vazia se não encontrado.
func CheckIP(ip string, networks []NetworkInfo) string {
    // Converte a string IP para net.IP
    parsedIP := net.ParseIP(ip)
    if parsedIP == nil {
        return "" // Retorna string vazia se o IP for inválido
    }

    // Verifica cada rede na lista
    for _, network := range networks {
        // Converte a string de rede para *net.IPNet
        _, ipNet, err := net.ParseCIDR(network.Network)
        if err != nil {
            continue // Ignora redes mal formatadas
        }
        // Verifica se o IP está contido na rede
        if ipNet.Contains(parsedIP) {
            return network.Info // Retorna a informação associada à rede
        }
    }

    return "" // Retorna string vazia se o IP não estiver em nenhuma rede
}