package cheat

import (
	"github.com/kbinani/win"
	"github.com/restartfu/emp/emp"
	"strings"
	"sync"
)

var (
	cheats   = map[string]*Cheat{}
	cheatsMu sync.Mutex
)

// Register registers a new cheat.
func Register(cheat *Cheat) {
	cheatsMu.Lock()
	cheats[cheat.Name()] = cheat
	cheatsMu.Unlock()
}

// ByName returns a cheat by name.
func ByName(name string) *Cheat {
	cheatsMu.Lock()
	defer cheatsMu.Unlock()
	return cheats[name]
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
	address   win.LPVOID

	defaultValue      float32
	name, description string
}

// New creates a new cheat.
func New(handle *emp.Handler, address uintptr, name, description string, defaultValue float32, updatable ...Updatable) *Cheat {
	return &Cheat{
		updatable:    updatable,
		handle:       handle,
		address:      win.LPVOID(address),
		defaultValue: defaultValue,
		name:         strings.ToLower(name),
		description:  description,
	}
}

// Address returns the address of the cheat.
func (c *Cheat) Address() win.LPVOID {
	return c.address
}

// Name returns the name of the cheat.
func (c *Cheat) Name() string {
	return c.name
}

// Description returns the description of the cheat.
func (c *Cheat) Description() string {
	return c.description
}

// Update enables the cheat.
func (c *Cheat) Update() {
	for _, t := range c.updatable {
		t.Update(c.handle, c.address)
	}
}

// DefaultValue returns the default value of the cheat.
func (c *Cheat) DefaultValue() float32 {
	return c.defaultValue
}

// SetValue sets the value of the cheat.
func (c *Cheat) SetValue(value float32) {
	for _, t := range c.updatable {
		t.SetValue(value)
	}
}
