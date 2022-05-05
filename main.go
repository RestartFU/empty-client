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
)

func main() {
	h := emp.New()
	for _, c := range []*cheat.Cheat{
		cheat.New("Reach", "Hurt entities from further away.", h, 3, cheat.NewReach(win.LPVOID(h.GameID()+0x440C8E0))),
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
				fmt.Println(color.HiCyanString(" list of available cheats:"))
				for _, c := range cheat.All() {
					fmt.Println(color.CyanString("  %s - %s", c.DisplayName(), c.Description()))
				}
				fmt.Print(color.CyanString(">: "))
			default:
				c, ok := cheat.ByName(text)
				if !ok {
					fmt.Println(color.CyanString("invalid cheat name"))
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
	fmt.Println(color.HiCyanString(`
Type help to see the list of available cheats, and type exit to gracefully exit the program.
 _____           _       
|   __|_____ ___| |_ _ _ 
|   __|     | . |  _| | |
|_____|_|_|_|  _|_| |_  |
            |_|     |___|
`))
}
