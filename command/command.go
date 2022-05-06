package command

import "sync"

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

	name, description string
}

// New creates a new Command.
func New(name, description string, runnables ...Runnable) *Command {
	return &Command{
		name:        name,
		description: description,
		runnables:   runnables,
	}
}

// Name returns the name of the command.
func (cmd *Command) Name() string {
	return cmd.name
}

// Description returns the description of the command.
func (cmd *Command) Description() string {
	return cmd.description
}

// Execute executes the command.
func (cmd *Command) Execute() {
	for _, runnable := range cmd.runnables {
		runnable.Run()
	}
}
