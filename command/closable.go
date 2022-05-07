package command

import "github.com/restartfu/emp/empty"

type closeable interface {
	close(*empty.Handler)
}
