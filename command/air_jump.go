package command

import (
	"errors"
	"github.com/kbinani/win"
	"github.com/restartfu/empty-client/empty"
	"time"
	"unsafe"
)

var airJumpValue float32 = 2.451060728e-38

type AirJump struct {
	enabled bool
}

// Run ...
func (a AirJump) Run(h *empty.Handler, args ...string) error {
	if arg := args[0]; arg == "enable" {
		if a.enabled {
			return errors.New("air jump is already enabled")
		}
		a.enabled = true
		go func() {
			for a.enabled {
				var num win.DWORD
				var bytesWritten win.SIZE_T
				address := win.LPVOID(h.LocalPlayer() + 0x000001D8)

				win.VirtualProtectEx(h.Handle(), address, 4, 0x40, &num)
				win.WriteProcessMemory(h.Handle(), address, uintptr(unsafe.Pointer(&airJumpValue)), 4, &bytesWritten)
				time.Sleep(time.Millisecond)
			}
		}()
	} else if arg == "disable" {
		a.enabled = false
		return nil
	} else {
		return errors.New("airjump argument must be either enable or disable")
	}
	return nil
}

// Name ...
func (a *AirJump) Name() string {
	return "Air Jump"
}

// Description ...
func (a *AirJump) Description() string {
	return ""
}

// HasInput ...
func (a *AirJump) HasInput() bool {
	return true
}
