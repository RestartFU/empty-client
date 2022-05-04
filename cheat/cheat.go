package cheat

import (
	"github.com/restartfu/emp/emp"
)

var cheats []*Cheat

func Register(cheat *Cheat) {
	cheats = append(cheats, cheat)
}

// Cheat contains the information about a cheat and its toggleable.
type Cheat struct {
	updatable []Updatable
	handle    *emp.Handler

	name        string
	description string
}

// New creates a new cheat.
func New(name, description string, handle *emp.Handler, updatable ...Updatable) *Cheat {
	return &Cheat{
		updatable:   updatable,
		handle:      handle,
		name:        name,
		description: description,
	}
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
		t.Update(c.handle)
	}
}

func (c *Cheat) SetValue(value float32) {
	for _, t := range c.updatable {
		t.SetValue(value)
	}
}
