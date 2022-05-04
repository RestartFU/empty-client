package main

import (
	"fmt"
	"github.com/TheTitanrain/w32"
	"github.com/fatih/color"
	"github.com/kbinani/win"
	"github.com/restartfu/emp/cheat"
	"github.com/restartfu/emp/emp"
	"syscall"
	"unsafe"
)

func main() {
	h := emp.New()
	for _, c := range []*cheat.Cheat{
		cheat.New("Reach", "Hurt entities from further away.", h, 3, cheat.NewReach(win.LPVOID(h.GameID()+0x440C8E0))),
	} {
		cheat.Register(c)
	}

	var actualFov float32
	readProcessMemory(w32.HANDLE(h.Handle()), 0x1809E9EF098, uintptr(unsafe.Pointer(&actualFov)), unsafe.Sizeof(actualFov))
	zoom := cheat.NewZoom(0x1809E9EF098)

	go func() {
		var pressed bool
		for {
			keyState := w32.GetAsyncKeyState(70)
			if keyState&(1<<15) != 0 {
				pressed = true
			} else {
				pressed = false
			}
			if !pressed {
				zoom.SetValue(actualFov)
			} else {
				zoom.SetValue(10)
			}
			zoom.Update(h)
		}
	}()

	go func() {
		<-h.Close
		for _, c := range cheat.All() {
			c.SetValue(c.DefaultValue())
			c.Update()
		}
		zoom.SetValue(actualFov)
		zoom.Update(h)
	}()
	fmt.Println(color.HiCyanString(`
Type help to see the list of available cheats, and type exit to gracefully exit the program.
 _____           _       
|   __|_____ ___| |_ _ _ 
|   __|     | . |  _| | |
|_____|_|_|_|  _|_| |_  |
            |_|     |___|
`))
	for {
		var v string
		fmt.Print(color.CyanString(">: "))
		_, _ = fmt.Scan(&v)
		switch v {
		case "exit":
			for _, c := range cheat.All() {
				c.SetValue(c.DefaultValue())
				c.Update()
			}
			return
		case "help":
			fmt.Println(color.HiCyanString("   Available cheats:"))
			for _, c := range cheat.All() {
				fmt.Println(color.CyanString("    %s: %s", c.DisplayName(), c.Description()))
			}
			continue
		}
		c, ok := cheat.ByName(v)
		if !ok {
			fmt.Println(color.CyanString("Invalid cheat name"))
			continue
		}
		fmt.Print(color.CyanString("enter %s value >: ", c.Name()))
		c.SetValue(scanFloat32())
		c.Update()
	}
}

func scanFloat32() float32 {
	var v float32
	_, err := fmt.Scan(&v)
	for err != nil {
		fmt.Println(color.CyanString("Invalid value"))
		_, err = fmt.Scan(&v)
	}
	return v
}

var (
	modkernel32           = syscall.NewLazyDLL("kernel32.dll")
	procReadProcessMemory = modkernel32.NewProc("ReadProcessMemory")
)

func readProcessMemory(hProcess w32.HANDLE, lpBaseAddress, lpBuffer, nSize uintptr) (lpNumberOfBytesRead int, ok bool) {
	var nBytesRead int
	ret, _, _ := procReadProcessMemory.Call(
		uintptr(hProcess),
		lpBaseAddress,
		lpBuffer,
		nSize,
		uintptr(unsafe.Pointer(&nBytesRead)),
	)

	return nBytesRead, ret != 0
}
