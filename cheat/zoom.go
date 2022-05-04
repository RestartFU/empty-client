package cheat

import (
	"github.com/kbinani/win"
	"github.com/restartfu/emp/emp"
	"unsafe"
)

// Zoom is a cheat used change your fov.
type Zoom struct {
	value float32

	pointer win.LPVOID
}

// NewZoom creates a new reach cheat.
func NewZoom(pointer win.LPVOID) *Zoom {
	return &Zoom{
		pointer: pointer,
	}
}

// Update ...
func (z *Zoom) Update(h *emp.Handler) {
	var bytesWritten win.SIZE_T

	win.WriteProcessMemory(h.Handle(), z.pointer, uintptr(unsafe.Pointer(&z.value)), 4, &bytesWritten)
}

// SetValue ...
func (z *Zoom) SetValue(value float32) {
	z.value = value
}
