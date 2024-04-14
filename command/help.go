package command

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/restartfu/empty-client/empty"
	"strings"
)

type Help struct{}

func (h Help) Run(*empty.Handler, ...string) error {
	for _, c := range All() {
		fmt.Println(color.YellowString("%s - %s", strings.ToUpper(c.Name()), c.Description()))
	}
	fmt.Println()
	return nil
}
func (h Help) Name() string {
	return "help"
}
func (h Help) Description() string {
	return ""
}
func (h Help) HasInput() bool {
	return false
}
