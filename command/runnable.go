package command

import "github.com/restartfu/empty-client/empty"

type Runnable interface {
	Run(*empty.Handler, ...string) error
	Name() string
	Description() string
	HasInput() bool
}
