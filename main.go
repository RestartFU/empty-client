package main

import (
	"fmt"
	"github.com/TheTitanrain/w32"
	"github.com/fatih/color"
	"github.com/kbinani/win"
	"github.com/restartfu/emp/command"
	"github.com/restartfu/emp/empty"
	"golang.org/x/sys/windows"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	c = make(chan os.Signal)
)

func main() {
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	go scan(handler())
	<-c
	for _, c := range command.All() {
		c.Close()
	}

}
func scan(h *empty.Handler) {
	for {
		fmt.Print(color.CyanString("|>: "))
		if cmd := command.ByName(scanInput()); cmd != nil {
			if len(cmd.Runnables()) > 1 {
				var formattedRunnables []string
				for n, c := range cmd.Runnables() {
					formattedRunnables = append(formattedRunnables, color.CyanString("[%v] %s - %s", n+1, c.Name(), c.Description()))
				}
				fmt.Println(strings.Join(formattedRunnables, "\n") + "\n")
			}
			r := scanRunnable(cmd)
			if !r.HasInput() {
				_ = r.Run(h)
				continue
			}
			fmt.Print(color.CyanString("- %s |>: ", strings.ToUpper(r.Name())))
			err := r.Run(h, strings.Split(scanInput(), " ")...)
			for err != nil {
				color.Red("/~\\ %s /~\\", err.Error())
				fmt.Print(color.CyanString("- %s |>: ", strings.ToUpper(r.Name())))
				err = r.Run(h, strings.Split(scanInput(), " ")...)
			}

			continue
		}
		color.Red("/~\\ unknown command /~\\")
	}
}

func scanIndex(cmd *command.Command) int {
	scan := func() (int, error) {
		fmt.Print(color.CyanString("- %s |>: ", strings.ToUpper(cmd.Name())))
		input := scanInput()
		return strconv.Atoi(input)
	}
	index, err := scan()
	for err != nil {
		color.Red("/~\\ invalid index /~\\")
		index, err = scan()
	}
	return index
}

func scanRunnable(cmd *command.Command) command.Runnable {
	if len(cmd.Runnables()) == 1 {
		return cmd.Runnables()[0]
	}
	scan := func() (command.Runnable, error) {
		r, err := cmd.Runnable(scanIndex(cmd) - 1)
		return r, err
	}
	r, err := scan()
	for err != nil {
		color.Red("/~\\ %s /~\\", err.Error())
		r, err = scan()
	}
	return r
}

func scanInput() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil && err.Error() == "EOF" {
		os.Exit(0)
	}
	return input
}

func handler() *empty.Handler {
	open := func() bool {
		return w32.FindWindowW(nil, windows.StringToUTF16Ptr("Minecraft")) != 0
	}

	for !open() {
		color.Cyan("Game window not found. Please open minecraft\n")
		for !open() {
			time.Sleep(time.Second)
		}
		time.Sleep(time.Second * 5)
	}
	h := empty.New()
	welcome(h)
	return h
}

func registerCommands(h *empty.Handler) {
	for _, c := range []*command.Command{
		command.New(h, "help", "See the list of available commands and cheats.", command.Help{}),
		command.New(h, "reach", "Hurt entities from a distance.", command.Reach{}),
		command.New(h, "airjump", "Jump, even while being in the air", &command.AirJump{}),
	} {
		command.Register(c)
	}
}

func welcome(h *empty.Handler) {
	win.SetConsoleTitle("Empty - By RestartFU <3")
	win.SetConsoleIcon(win.HICON(h.Handle()))

	registerCommands(h)
	fmt.Println(color.CyanString(`Empty v0.1
 _____           _       
|   __|_____ ___| |_ _ _ 
|   __|     | . |  _| | |
|_____|_|_|_|  _|_| |_  |
            |_|     |___|

Type 'help' to see the list of available commands and cheats.
`))
}
