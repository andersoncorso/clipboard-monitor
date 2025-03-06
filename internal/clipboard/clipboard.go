package clipboard

import (
	"fmt"

	"golang.design/x/clipboard"
)

// GetContent recupera o conteúdo de texto da área de transferência do sistema.
// Retorna uma string com o conteúdo e um erro, se houver.
func GetContent() (string, error) {
    // Inicializa o sistema de clipboard.
    // Isso é necessário antes de qualquer operação com a área de transferência.
    err := clipboard.Init()
    if err != nil {
        // Se a inicialização falhar, retorna um erro formatado.
        return "", fmt.Errorf("falha ao inicializar a área de transferência: %w", err)
    }

    // Verifica se há texto disponível na área de transferência.
    // A função Read retorna um slice de bytes, então verificamos se o comprimento é zero.
    if len(clipboard.Read(clipboard.FmtText)) == 0 {
        // Se não houver texto, retorna uma string vazia sem erro.
        // Isso permite que o chamador diferencie entre "sem texto" e "erro".
        return "", nil
    }

    // Lê o conteúdo de texto da área de transferência.
    content := clipboard.Read(clipboard.FmtText)
    
    // Converte o slice de bytes para string e retorna.
    return string(content), nil
}