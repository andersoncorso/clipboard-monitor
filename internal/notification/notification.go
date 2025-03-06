package notification

import (
	"github.com/gen2brain/beeep"
)

// Show exibe uma notificação do sistema com o título e mensagem fornecidos.
// Esta função é uma wrapper simples para a função Notify da biblioteca beeep.
//
// Parâmetros:
//   - title: O título da notificação.
//   - message: O corpo da mensagem da notificação.
//
// Retorna:
//   - error: Um erro, se houver algum problema ao exibir a notificação.
func Show(title, message string) error {
    // A função beeep.Notify é chamada com os parâmetros fornecidos.
    // O terceiro parâmetro (caminho do ícone) é deixado vazio ("").
    return beeep.Notify(title, message, "")
}