package main

import (
	"fmt"
	"github.com/TheTitanrain/w32"
	"github.com/fatih/color"
	"github.com/kbinani/win"
	"golang.org/x/sys/windows"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	processId := getGameProcessId()
	gameWindow := w32.FindWindowW(nil, windows.StringToUTF16Ptr("Minecraft"))
	if gameWindow == 0 {
		color.Cyan("Game window not found. Please open minecraft\n")
		for gameWindow == 0 {
			gameWindow = w32.FindWindowW(nil, windows.StringToUTF16Ptr("Minecraft"))
		}
		time.Sleep(time.Second * 5)
	}

	h := win.OpenProcess(w32.PROCESS_ALL_ACCESS, true, win.DWORD(processId))

	var reach float32
	reachPtr := win.LPVOID(0x7FF67B6DC8E0)

	fmt.Print(color.CyanString("Enter reach: "))
	_, err := fmt.Scan(&reach)
	if err != nil {
		return
	}

	var num win.DWORD = 0
	var bytesWritten win.SIZE_T

	win.VirtualProtectEx(h, reachPtr, 4, 0x40, &num)
	win.WriteProcessMemory(h, reachPtr, uintptr(unsafe.Pointer(&reach)), 4, &bytesWritten)
	for {
	}
}

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
