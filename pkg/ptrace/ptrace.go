package ptrace

import (
	"runtime"
)

var ch chan func()

func init() {
	ch = make(chan func())

	go func() {
		runtime.LockOSThread()
		for f := range ch {
			f()
		}
	}()
}

func Do(f func() error) error {
	errch := make(chan error, 1)
	ch <- func() { errch <- f() }
	return <-errch
}
