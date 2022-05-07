package command

import "github.com/restartfu/empty-client/empty"

type closeable interface {
	close(*empty.Handler)
}
