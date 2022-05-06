package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/kbinani/win"
	"github.com/restartfu/emp/cheat"
	"github.com/restartfu/emp/command"
	"github.com/restartfu/emp/emp"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	h := emp.New()
	win.SetConsoleIcon(win.HICON(h.Handle()))

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
				var exit = errors.New("exit")
				f := func() error {
					var v any
					fmt.Print(color.CyanString("- %s (DEFAULT: %v) |>: ", strings.ToUpper(cht.Name()), cht.DefaultValue()))

					_, err := fmt.Scan(&input)
					if err != nil && err == io.EOF {
						os.Exit(0)
					}
					if input == "exit" {
						return exit
					}

					v, err = strconv.ParseFloat(input, 32)
					if err != nil {
						v = nil
					}

					err = cht.SetValue(v)
					if err != nil {
						color.Red("/~\\ %s /~\\", err.Error())
					}
					return err

				}
				for err := f(); err != nil && err != exit; err = f() {
				}
				if err != exit {
					cht.Update()
				}
				continue
			}
			color.Red("/~\\ unknown command or cheat /~\\")
		}
	}()
	<-h.Close
	for _, c := range cheat.All() {
		_ = c.SetValue(c.DefaultValue())
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
	win.SetConsoleTitle("Empty - By RestartFU <3")
	fmt.Println(color.CyanString(`Empty v0.1
 _____           _       
|   __|_____ ___| |_ _ _ 
|   __|     | . |  _| | |
|_____|_|_|_|  _|_| |_  |
            |_|     |___|

Type 'help' to see the list of available commands and cheats.
`))
}
