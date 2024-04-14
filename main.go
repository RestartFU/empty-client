package main

import (
	"bufio"
	"fmt"
	"github.com/TheTitanrain/w32"
	"github.com/kbinani/win"
	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"github.com/restartfu/empty-client/command"
	"github.com/restartfu/empty-client/empty"
	"golang.org/x/sys/windows"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func init() {
	log.SetOutput(colorable.NewColorableStdout())

}

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
		fmt.Print(aurora.Yellow("-> "))
		if cmd := command.ByName(scanInput()); cmd != nil {
			if len(cmd.Runnables()) > 1 {
				var formattedRunnables []string
				for n, c := range cmd.Runnables() {
					formattedRunnables = append(formattedRunnables, fmt.Sprintf("[%v] %s - %s", n+1, c.Name(), c.Description()))
				}
				fmt.Println(strings.Join(formattedRunnables, "\n") + "\n")
			}
			r := scanRunnable(cmd)
			if !r.HasInput() {
				_ = r.Run(h)
				continue
			}
			fmt.Print(aurora.Yellow(fmt.Sprintf("%s -> ", strings.ToUpper(r.Name()))))
			err := r.Run(h, strings.Split(scanInput(), " ")...)
			for err != nil {
				fmt.Println(aurora.Red(fmt.Sprintf("%s", err.Error())))
				fmt.Print(aurora.Yellow(fmt.Sprintf("%s -> ", strings.ToUpper(r.Name()))))
				err = r.Run(h, strings.Split(scanInput(), " ")...)
			}

			continue
		}
		fmt.Println(aurora.Red("Unknown command, try help"))
	}
}

func scanIndex(cmd *command.Command) int {
	scan := func() (int, error) {
		fmt.Print(aurora.Yellow(fmt.Sprintf("%s -> ", strings.ToUpper(cmd.Name()))))
		input := scanInput()
		return strconv.Atoi(input)
	}
	index, err := scan()
	for err != nil {
		fmt.Println(aurora.Red("Invalid prompt"))
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
		fmt.Println(aurora.Yellow(fmt.Sprintf("[ERROR] %s", err.Error())))
		r, err = scan()
	}
	return r
}

func scanInput() string {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	for strings.TrimSpace(s.Text()) == "" {
		s.Scan()
	}
	return strings.Split(strings.TrimSpace(s.Text()), " ")[0]
}

func handler() *empty.Handler {
	open := func() bool {
		return w32.FindWindowW(nil, windows.StringToUTF16Ptr("Minecraft")) != 0
	}

	for !open() {
		fmt.Println(aurora.Yellow("Game window not found. Please open minecraft\n"))
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
		command.New(h, "timer", "Increase the speed of your games.", command.Timer{}),
	} {
		command.Register(c)
	}
}

func welcome(h *empty.Handler) {
	win.SetConsoleTitle("Empty - By RestartFU <3")
	win.SetConsoleIcon(win.HICON(h.Handle()))

	registerCommands(h)
	fmt.Println(aurora.Yellow(`Empty v0.2
███████╗███╗   ███╗██████╗ ████████╗██╗   ██╗
██╔════╝████╗ ████║██╔══██╗╚══██╔══╝╚██╗ ██╔╝
█████╗  ██╔████╔██║██████╔╝   ██║    ╚████╔╝ 
██╔══╝  ██║╚██╔╝██║██╔═══╝    ██║     ╚██╔╝  
███████╗██║ ╚═╝ ██║██║        ██║      ██║   
╚══════╝╚═╝     ╚═╝╚═╝        ╚═╝      ╚═╝   

Type 'help' to see the list of available commands and cheats.
`))
}
