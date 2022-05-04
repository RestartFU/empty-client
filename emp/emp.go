package emp

import (
	"github.com/TheTitanrain/w32"
	"github.com/fatih/color"
	"github.com/kbinani/win"
	"golang.org/x/sys/windows"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"
)

// Handler is a handler which contains everything you need to know about a process
type Handler struct {
	h         win.HANDLE
	processID uint32
	gameID    uintptr

	Close chan os.Signal
}

// New creates a new emp handler
func New() *Handler {
	gameWindow := w32.FindWindowW(nil, windows.StringToUTF16Ptr("Minecraft"))
	if gameWindow == 0 {
		color.Cyan("Game window not found. Please open minecraft\n")
		for gameWindow == 0 {
			gameWindow = w32.FindWindowW(nil, windows.StringToUTF16Ptr("Minecraft"))
		}
		time.Sleep(time.Second * 5)
	}

	processID := getGameProcessId()
	h := &Handler{
		h:         win.OpenProcess(w32.PROCESS_ALL_ACCESS, true, win.DWORD(processID)),
		processID: processID,
		gameID:    getGameModule(processID),
		Close:     make(chan os.Signal),
	}
	signal.Notify(h.Close, os.Interrupt, syscall.SIGTERM)
	return h
}

// Handle returns the handle to the game process
func (h *Handler) Handle() win.HANDLE {
	return h.h
}

// ProcessID returns the process ID of the game process
func (h *Handler) ProcessID() uint32 {
	return h.processID
}

// GameID returns the game's module ID
func (h *Handler) GameID() uintptr {
	return h.gameID
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
