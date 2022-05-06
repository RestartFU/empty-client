package cheat

import (
	"github.com/kbinani/win"
	"github.com/restartfu/emp/emp"
	"time"
	"unsafe"
)

var AirJumpValue float32 = 2.451060728e-38

type AirJump struct {
	value float32

	pointer win.LPVOID
}

// NewAirJump creates a new cheat that enables double jump.
func NewAirJump(pointer win.LPVOID) *AirJump {
	return &AirJump{
		pointer: pointer,
	}
}

// Update ...
func (d *AirJump) Update(h *emp.Handler) {
	go func() {
		for d.value == 1 {
			var num win.DWORD
			var bytesWritten win.SIZE_T
			win.VirtualProtectEx(h.Handle(), d.pointer, 4, 0x40, &num)
			win.WriteProcessMemory(h.Handle(), d.pointer, uintptr(unsafe.Pointer(&AirJumpValue)), 4, &bytesWritten)
			time.Sleep(time.Millisecond)
		}
	}()
}

// SetValue ...
func (d *AirJump) SetValue(value float32) {
	d.value = value
}
