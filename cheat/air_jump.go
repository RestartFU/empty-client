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
}

// Update ...
func (d *AirJump) Update(h *emp.Handler, address win.LPVOID) {
	go func() {
		for d.value == 1 {
			var num win.DWORD
			var bytesWritten win.SIZE_T
			win.VirtualProtectEx(h.Handle(), address, 4, 0x40, &num)
			win.WriteProcessMemory(h.Handle(), address, uintptr(unsafe.Pointer(&AirJumpValue)), 4, &bytesWritten)
			time.Sleep(time.Millisecond)
		}
	}()
}

// SetValue ...
func (d *AirJump) SetValue(value float32) {
	d.value = value
}
