package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/kbinani/win"
	"github.com/restartfu/emp/cheat"
	"github.com/restartfu/emp/emp"
)

func main() {
	h := emp.New()
	r := cheat.New("Reach", "", h, cheat.NewReach(win.LPVOID(h.GameID()+0x440C8E0)))

	var reach float32
	fmt.Print(color.CyanString("Enter reach: "))
	_, err := fmt.Scan(&reach)
	if err != nil {
		return
	}
	r.SetValue(reach)
	r.Update()
}
