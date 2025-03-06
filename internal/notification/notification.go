package notification

import (
	"github.com/gen2brain/beeep"
)

func Show(title, message string) error {
    return beeep.Notify(title, message, "")
}