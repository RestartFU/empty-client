package empty

import (
	"github.com/TheTitanrain/w32"
	"github.com/kbinani/win"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// Handler is a handler which contains everything you need to know about a process.
type Handler struct {
	h win.HANDLE

	processID  uint32
	gameID     uintptr
	gameWindow win.HWND
}

// New creates a new empty handler.
func New() *Handler {
	processID := getGameProcessId()
	h := &Handler{
		h:          win.OpenProcess(w32.PROCESS_ALL_ACCESS, true, win.DWORD(processID)),
		processID:  processID,
		gameID:     getGameModule(processID),
		gameWindow: win.HWND(w32.FindWindowW(nil, windows.StringToUTF16Ptr("Minecraft"))),
	}
	return h
}

// Handle returns the handle to the game process.
func (h *Handler) Handle() win.HANDLE {
	return h.h
}

// ProcessID returns the process ID of the game process.
func (h *Handler) ProcessID() uint32 {
	return h.processID
}

// GameID returns the game's module ID.
func (h *Handler) GameID() uintptr {
	return h.gameID
}

// GameWindow returns the game window.
func (h *Handler) GameWindow() win.HWND {
	return h.gameWindow
}

// FindAddressOffset finds returns an address with the given offset.
func (h *Handler) FindAddressOffset(providedPtr uintptr, providedOffsets []uintptr) uintptr {
	address := providedPtr
	for _, offset := range providedOffsets {
		win.ReadProcessMemory(h.Handle(), address, win.LPVOID(unsafe.Pointer(&address)), win.SIZE_T(unsafe.Sizeof(address)), nil)
		address += offset
	}
	return address
}

// Focused returns true if the game is focused.
func (h *Handler) Focused() bool {
	return win.GetForegroundWindow() == h.GameWindow()
}

// getGameProcessId returns the process ID of the game process
func getGameProcessId() uint32 {
	var processID uint32
	snapshotHandle, err := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	defer syscall.Close(snapshotHandle)
	if err == nil {
		var processEntry syscall.ProcessEntry32
		processEntry.Size = uint32(unsafe.Sizeof(processEntry))
		if syscall.Process32First(snapshotHandle, &processEntry) == nil {
			for {
				if windows.UTF16ToString(processEntry.ExeFile[:]) == "Minecraft.Windows.exe" {
					processID = processEntry.ProcessID
					break
				}

				if err := syscall.Process32Next(snapshotHandle, &processEntry); err != nil {
					break
				}
			}
		}
	} else {
		panic(err)
	}

	return processID
}

// getGameModule returns the game's module ID
func getGameModule(processID uint32) uintptr {
	var gameModuleAddress uintptr = 0
	snapshotHandle := w32.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPMODULE|syscall.TH32CS_SNAPMODULE32, processID)
	defer w32.CloseHandle(snapshotHandle)

	if snapshotHandle != w32.ERROR_INVALID_HANDLE {
		var modEntry w32.MODULEENTRY32
		modEntry.Size = uint32(unsafe.Sizeof(modEntry))

		if w32.Module32First(snapshotHandle, &modEntry) {
			for {
				if windows.UTF16ToString(modEntry.SzModule[:]) == "Minecraft.Windows.exe" {
					gameModuleAddress = uintptr(unsafe.Pointer(modEntry.ModBaseAddr))
					break
				}

				if w32.Module32Next(snapshotHandle, &modEntry) {
					break
				}
			}
		}
	}

	return gameModuleAddress
}
