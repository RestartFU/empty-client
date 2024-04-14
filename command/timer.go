package command

import (
	"github.com/kbinani/win"
	"github.com/pkg/errors"
	"github.com/restartfu/empty-client/empty"
	"strconv"
	"unsafe"
)

// Timer is a cheat used to accelerate game.
type Timer struct{}

// Run ...
func (r Timer) Run(h *empty.Handler, args ...string) error {
	v, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return errors.New("Timer value must be a float.")
	}
	value := v * 1000
	r.PatchBytesFloat32(h, win.LPVOID(h.GameID()+0x4DE1A58), value, win.SIZE_T(unsafe.Sizeof(value)))
	return nil
}

func (r Timer) PatchBytesFloat32(h *empty.Handler, dst win.LPVOID, value float64, size win.SIZE_T) {
	var oldprotect win.DWORD
	var bytesWritten win.SIZE_T

	win.VirtualProtectEx(h.Handle(), dst, size, 0x40, &oldprotect)
	win.WriteProcessMemory(h.Handle(), dst, uintptr(unsafe.Pointer(&value)), size, &bytesWritten)
	win.VirtualProtectEx(h.Handle(), dst, size, oldprotect, &oldprotect)
}

func (r Timer) Name() string {
	return "Timer"
}
func (r Timer) Description() string {
	return ""
}
func (r Timer) HasInput() bool {
	return true
}

func (r *Timer) Close(h *empty.Handler) {
	_ = r.Run(h, "3")
}
