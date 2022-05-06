package cheat

import (
	"github.com/kbinani/win"
	"github.com/restartfu/emp/emp"
)

// Updatable is an interface with methods for a cheat that can be updated.
type Updatable interface {
	Update(*emp.Handler, win.LPVOID)
	SetValue(float32)
}
