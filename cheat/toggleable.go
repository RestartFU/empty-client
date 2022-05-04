package cheat

import (
	"github.com/restartfu/emp/emp"
)

type Toggleable interface {
	Update(*emp.Handler)

	SetValue(float32)
}
