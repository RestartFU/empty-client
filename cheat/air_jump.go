package cheat

import (
	"fmt"
	"github.com/kbinani/win"
	"github.com/restartfu/emp/emp"
	"time"
	"unsafe"
)

var AirJumpValue float32 = 2.451060728e-38

type AirJump struct {
	enabled bool
}

// Update ...
func (d *AirJump) Update(h *emp.Handler, address win.LPVOID) {
	go func() {
		for d.enabled {
			var num win.DWORD
			var bytesWritten win.SIZE_T
			win.VirtualProtectEx(h.Handle(), address, 4, 0x40, &num)
			win.WriteProcessMemory(h.Handle(), address, uintptr(unsafe.Pointer(&AirJumpValue)), 4, &bytesWritten)
			time.Sleep(time.Millisecond)
		}
	}()
}

// SetValue ...
func (d *AirJump) SetValue(value any) error {
	v, ok := value.(float64)
	if !ok || v > 1 || v < 0 {
		return fmt.Errorf("value must be either 0, or 1. (1: enabled; 0: disabled)")
	}
	d.enabled = v == 1
	return nil
}
