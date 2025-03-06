package ipchecker

import (
	"bufio"
	"net"
	"os"
	"strings"
)

type NetworkInfo struct {
    Network string
    Info    string
}

func LoadNetworkData(filename string) ([]NetworkInfo, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var networks []NetworkInfo
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.SplitN(line, "|", 2)
        if len(parts) == 2 {
            networks = append(networks, NetworkInfo{
                Network: strings.TrimSpace(parts[0]),
                Info:    strings.TrimSpace(parts[1]),
            })
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return networks, nil
}

func CheckIP(ip string, networks []NetworkInfo) string {
    parsedIP := net.ParseIP(ip)
    if parsedIP == nil {
        return ""
    }

    for _, network := range networks {
        _, ipNet, err := net.ParseCIDR(network.Network)
        if err != nil {
            continue
        }
        if ipNet.Contains(parsedIP) {
            return network.Info
        }
    }

    return ""
}