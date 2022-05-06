package cheat

import (
	"github.com/kbinani/win"
	"github.com/restartfu/emp/emp"
	"unsafe"
)

// Reach is a cheat used to hit players from further away.
type Reach struct {
	value float32
}

// Update ...
func (r *Reach) Update(h *emp.Handler, address win.LPVOID) {
	var num win.DWORD
	var bytesWritten win.SIZE_T

	win.VirtualProtectEx(h.Handle(), address, 4, 0x40, &num)
	win.WriteProcessMemory(h.Handle(), address, uintptr(unsafe.Pointer(&r.value)), 4, &bytesWritten)
}

// SetValue ...
func (r *Reach) SetValue(value float32) {
	r.value = value
}
