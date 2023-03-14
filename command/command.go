package command

import (
	"fmt"
	"github.com/restartfu/empty-client/empty"
	"sync"
)

var (
	commands   = map[string]*Command{}
	commandsMu sync.Mutex
)

// Register registers a command.
func Register(cmd *Command) {
	commandsMu.Lock()
	defer commandsMu.Unlock()
	commands[cmd.name] = cmd
}

// ByName returns the Command with the given name.
func ByName(name string) *Command {
	commandsMu.Lock()
	defer commandsMu.Unlock()
	return commands[name]
}

// All returns a slice of all registered commands.
func All() (c []*Command) {
	commandsMu.Lock()
	defer commandsMu.Unlock()
	for _, cmd := range commands {
		c = append(c, cmd)
	}
	return c
}

// Command represents a command which can be executed by typing its name in the console.
type Command struct {
	runnables []Runnable
	handler   *empty.Handler
	input     bool

	name, description string
}

// New creates a new Command.
func New(handler *empty.Handler, name, description string, runnables ...Runnable) *Command {
	c := &Command{
		handler:     handler,
		name:        name,
		description: description,
	}
	for _, r := range runnables {
		c.runnables = append(c.runnables, r)
	}
	return c
}

// Name returns the name of the command.
func (cmd *Command) Name() string {
	return cmd.name
}

// Description returns the description of the command.
func (cmd *Command) Description() string {
	return cmd.description
}

// Runnables returns the runnables of the command.
func (cmd *Command) Runnables() []Runnable {
	return cmd.runnables
}

// Runnable returns the runnable with the given index.
func (cmd *Command) Runnable(index int) (Runnable, error) {
	if len(cmd.runnables)-1 < index || index < 0 {
		return nil, fmt.Errorf("unknown index: '%v'", index)
	}
	return cmd.runnables[index], nil
}

// Close closes the command.
func (cmd *Command) Close() {
	for _, r := range cmd.runnables {
		if c, ok := r.(closeable); ok {
			c.close(cmd.handler)
		}
	}
}
