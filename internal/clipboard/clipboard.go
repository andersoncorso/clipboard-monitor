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

    content := clipboard.Read(clipboard.FmtText)
    return string(content), nil
}