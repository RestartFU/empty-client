package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/kbinani/win"
	"github.com/restartfu/emp/cheat"
	"github.com/restartfu/emp/emp"
	"os"
)

func main() {
	h := emp.New()
	cheat.Register(cheat.New("Reach", "Hurt entities from further away.", h, 3, cheat.NewReach(win.LPVOID(h.GameID()+0x440C8E0))))

	go func() {
		<-h.Close
		for _, c := range cheat.All() {
			c.SetValue(c.DefaultValue())
			c.Update()
		}
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
			h.Close <- os.Interrupt
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
