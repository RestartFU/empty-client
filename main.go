package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/kbinani/win"
	"github.com/restartfu/emp/cheat"
	"github.com/restartfu/emp/command"
	"github.com/restartfu/emp/emp"
	"os"
	"strconv"
	"strings"
)

func main() {
	h := emp.New()

	registerCommands()
	registerCheats(h)
	welcome()

	go func() {
		for {
			var input string
			fmt.Print(color.CyanString("|>: "))
			_, err := fmt.Scan(&input)
			if err != nil && err.Error() == "EOF" {
				os.Exit(0)
			}
			if cmd := command.ByName(input); cmd != nil {
				cmd.Execute()
				continue
			}
			if cht := cheat.ByName(input); cht != nil {
				var value string
				var v float64
				f := func() error {
					fmt.Print(color.CyanString("- %s (DEFAULT: %v) |>: ", strings.ToUpper(cht.Name()), cht.DefaultValue()))
					_, err := fmt.Scan(&value)
					if err != nil {
						return err
					}
					v, err = strconv.ParseFloat(value, 32)
					return err
				}
				err = f()
				for err != nil {
					if err.Error() == "EOF" {
						os.Exit(0)
					}
					color.Cyan("Invalid value\n")
					err = f()
				}
				cht.SetValue(float32(v))
				cht.Update()
				continue
			}
			color.Cyan("Invalid command\n")
		}
	}()
	<-h.Close
	for _, c := range cheat.All() {
		c.SetValue(c.DefaultValue())
		c.Update()
	}
}

func registerCheats(h *emp.Handler) {
	for _, c := range []*cheat.Cheat{
		cheat.New(h, h.GameID()+0x440C8E0, "reach", "Hurt entities from further away.", 3, &cheat.Reach{}),
		cheat.New(h, h.LocalPlayer()+0x000001D8, "airjump", "Jump, even in the air. [1: enabled; 0: disabled]", cheat.AirJumpValue, &cheat.AirJump{}),
	} {
		cheat.Register(c)
	}
}

func registerCommands() {
	for _, c := range []*command.Command{
		command.New("help", "See the list of available commands and cheats.", command.Help{}),
	} {
		command.Register(c)
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

Type 'help' to see the list of available commands and cheats.
`))
}
