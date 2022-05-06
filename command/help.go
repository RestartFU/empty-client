package command

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/restartfu/emp/cheat"
	"strings"
)

type Help struct {
}

func (h Help) Run() {
	for _, c := range All() {
		fmt.Println(color.CyanString("%s - %s", strings.ToUpper(c.Name()), c.Description()))
	}
	fmt.Println()
	for _, c := range cheat.All() {
		fmt.Println(color.CyanString("%s - %s", strings.ToUpper(c.Name()), c.Description()))
	}
}
