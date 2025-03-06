package clipboard

import (
	"fmt"

	"golang.design/x/clipboard"
)

func GetContent() (string, error) {
    err := clipboard.Init()
    if err != nil {
        return "", fmt.Errorf("falha ao inicializar a área de transferência: %w", err)
    }

    // Verifica se há texto disponível na área de transferência
    if len(clipboard.Read(clipboard.FmtText)) == 0 {
        return "", nil // Retorna uma string vazia se não for texto
    }

    content := clipboard.Read(clipboard.FmtText)
    return string(content), nil
}