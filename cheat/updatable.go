package cheat

import (
	"github.com/restartfu/emp/emp"
)

type Updatable interface {
	Update(*emp.Handler)
	SetValue(float32)
}
