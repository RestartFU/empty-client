package cheat

import (
	"github.com/restartfu/emp/emp"
	"strings"
	"sync"
)

var (
	cheats   = map[string]*Cheat{}
	cheatsMu sync.Mutex
)

func Register(cheat *Cheat) {
	cheatsMu.Lock()
	cheats[cheat.Name()] = cheat
	cheatsMu.Unlock()
}

// ByName returns a cheat by name.
func ByName(name string) (*Cheat, bool) {
	cheatsMu.Lock()
	defer cheatsMu.Unlock()
	cheat, ok := cheats[name]
	return cheat, ok
}

// All returns all cheats
func All() (v []*Cheat) {
	cheatsMu.Lock()
	defer cheatsMu.Unlock()
	for _, c := range cheats {
		v = append(v, c)
	}
	return
}

// Cheat contains the information about a cheat and its toggleable.
type Cheat struct {
	updatable []Updatable
	handle    *emp.Handler

	defaultValue float32

	name        string
	displayName string
	description string
}

// New creates a new cheat.
func New(name, description string, handle *emp.Handler, defaultValue float32, updatable ...Updatable) *Cheat {
	return &Cheat{
		updatable:    updatable,
		handle:       handle,
		defaultValue: defaultValue,
		name:         strings.ToLower(name),
		displayName:  name,
		description:  description,
	}
}

// Name returns the name of the cheat.
func (c *Cheat) Name() string {
	return c.name
}
func (c *Cheat) DisplayName() string {
	return c.displayName
}

// Description returns the description of the cheat.
func (c *Cheat) Description() string {
	return c.description
}

// Update enables the cheat.
func (c *Cheat) Update() {
	for _, t := range c.updatable {
		t.Update(c.handle)
	}
}

func (c *Cheat) DefaultValue() float32 {
	return c.defaultValue
}

func (c *Cheat) SetValue(value float32) {
	for _, t := range c.updatable {
		t.SetValue(value)
	}
}
