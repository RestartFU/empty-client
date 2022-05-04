package cheat

import (
	"github.com/kbinani/win"
	"github.com/restartfu/emp/emp"
	"unsafe"
)

// Reach is a cheat used to hit players from further away.
type Reach struct {
	value float32

	pointer win.LPVOID
}

// NewReach creates a new reach cheat.
func NewReach(pointer win.LPVOID) *Reach {
	return &Reach{
		pointer: pointer,
	}
}

// Update ...
func (r *Reach) Update(h *emp.Handler) {
	var num win.DWORD
	var bytesWritten win.SIZE_T

	win.VirtualProtectEx(h.Handle(), r.pointer, 4, 0x40, &num)
	win.WriteProcessMemory(h.Handle(), r.pointer, uintptr(unsafe.Pointer(&r.value)), 4, &bytesWritten)
}

// SetValue ...
func (r *Reach) SetValue(value float32) {
	r.value = value
}
