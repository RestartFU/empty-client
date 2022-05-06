package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/kbinani/win"
	"github.com/restartfu/emp/cheat"
	"github.com/restartfu/emp/emp"
	"os"
	"strconv"
	"strings"
)

func main() {
	h := emp.New()
	localPlayer := h.FindAddressOffset(h.GameID()+0x0549E7F8, []uintptr{0x20, 0x0, 0x18, 0xB8, 0x198, 0x0, 0x0})
	for _, c := range []*cheat.Cheat{
		cheat.New("reach", "Hurt entities from further away.", h, 3, cheat.NewReach(win.LPVOID(h.GameID()+0x440C8E0))),
		cheat.New("airjump", "Jump, even in the air. [1: enabled; 0: disabled]", h, cheat.AirJumpValue, cheat.NewAirJump(win.LPVOID(localPlayer+0x000001D8))),
	} {
		cheat.Register(c)
	}
	welcome()
	go func() {
		var enteringValue bool
		var currentCheat *cheat.Cheat

		fmt.Print(color.CyanString(">: "))
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()

			if enteringValue {
				v, err := strconv.ParseFloat(text, 32)
				if err != nil {
					fmt.Println(color.CyanString("Invalid value"))
					continue
				}
				currentCheat.SetValue(float32(v))
				currentCheat.Update()

				enteringValue = false
				currentCheat = nil
				fmt.Print(color.CyanString(">: "))
				continue
			}
			switch text {
			case "exit":
				os.Exit(0)
			case "help":
				fmt.Println()
				for _, c := range cheat.All() {
					fmt.Println(color.CyanString("%s - %s", strings.ToUpper(c.Name()), c.Description()))
				}
				fmt.Print(color.CyanString("\n>: "))
			default:
				c, ok := cheat.ByName(text)
				if !ok {
					fmt.Println(color.CyanString("invalid cheat name"))
					fmt.Print(color.CyanString(">: "))
					continue
				}
				enteringValue = true
				currentCheat = c
				fmt.Print(color.CyanString("enter %s value >: ", currentCheat.Name()))
			}
		}
		h.Close <- os.Interrupt
	}()
	<-h.Close
	for _, c := range cheat.All() {
		c.SetValue(c.DefaultValue())
		c.Update()
	}
}

func welcome() {
	win.SetConsoleTitle("EMP")
	fmt.Println(color.CyanString(`Empty v0.1
 _____           _       
|   __|_____ ___| |_ _ _ 
|   __|     | . |  _| | |
|_____|_|_|_|  _|_| |_  |
            |_|     |___|
Type 'help' to see the list of available cheats.
`))
}
