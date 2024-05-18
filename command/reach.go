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

	survivalBytes := []byte{ 0xF3, 0x44, 0x0F, 0x10, 0x3D, 0x7D, 0x22, 0xAF, 0x04  }
	nop := []byte{0x90, 0x90}

	r.PatchBytes(h, win.LPVOID(h.GameID()+0x5CD7F8), nop, win.SIZE_T(len(nop)))
	r.PatchBytesFloat32(h, win.LPVOID(h.GameID()+0x50BFA80), value, win.SIZE_T(unsafe.Sizeof(value)))
	r.PatchBytes(h, win.LPVOID(h.GameID()+0x5CD7FA), survivalBytes, win.SIZE_T(len(survivalBytes)))
	return nil
}

func (r Reach) PatchBytesFloat32(h *empty.Handler, dst win.LPVOID, value float32, size win.SIZE_T) {
	var oldprotect win.DWORD
	var bytesWritten win.SIZE_T

	win.VirtualProtectEx(h.Handle(), dst, size, 0x40, &oldprotect)
	win.WriteProcessMemory(h.Handle(), dst, uintptr(unsafe.Pointer(&value)), size, &bytesWritten)
	win.VirtualProtectEx(h.Handle(), dst, size, oldprotect, &oldprotect)
}

func (r Reach) PatchBytes(h *empty.Handler, dst win.LPVOID, value []byte, size win.SIZE_T) {
	var oldprotect win.DWORD
	var bytesWritten win.SIZE_T

	win.VirtualProtectEx(h.Handle(), dst, size, 0x40, &oldprotect)
	win.WriteProcessMemory(h.Handle(), dst, uintptr(unsafe.Pointer(&value[0])), size, &bytesWritten)
	win.VirtualProtectEx(h.Handle(), dst, size, oldprotect, &oldprotect)
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
