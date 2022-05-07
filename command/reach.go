package command

import (
	"github.com/kbinani/win"
	"github.com/pkg/errors"
	"github.com/restartfu/empty-client/empty"
	"strconv"
	"unsafe"
)

// Reach is a cheat used to hit players from further away.
type Reach struct{}

// Run ...
func (r Reach) Run(h *empty.Handler, args ...string) error {
	v, err := strconv.ParseFloat(args[0], 32)
	if err != nil {
		return errors.New("reach value must be a float.")
	}
	value := float32(v)
	var num win.DWORD
	var bytesWritten win.SIZE_T
	address := win.LPVOID(h.GameID() + 0x440C8E0)

	win.VirtualProtectEx(h.Handle(), address, 4, 0x40, &num)
	win.WriteProcessMemory(h.Handle(), address, uintptr(unsafe.Pointer(&value)), 4, &bytesWritten)
	return nil
}

func (r Reach) Name() string {
	return "Reach"
}
func (r Reach) Description() string {
	return ""
}
func (r Reach) HasInput() bool {
	return true
}

func (r *Reach) Close(h *empty.Handler) {
	_ = r.Run(h, "3")
}
