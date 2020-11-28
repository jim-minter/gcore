package notes

// #include <sys/ptrace.h>
// #include <elf.h>
import "C"

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"

	"github.com/jim-minter/gcore/pkg/elf"
	"github.com/jim-minter/gcore/pkg/ptrace"
)

const X86_XSTATE_MAX_SIZE = 2696

func Xstate(pid, tid int) (*elf.Note, error) {
	var xstate [X86_XSTATE_MAX_SIZE]byte
	iov := unix.Iovec{Base: &xstate[0], Len: X86_XSTATE_MAX_SIZE}

	err := ptrace.Do(func() (err error) {
		_, _, errno := syscall.Syscall6(syscall.SYS_PTRACE, C.PTRACE_GETREGSET, uintptr(tid), C.NT_X86_XSTATE, uintptr(unsafe.Pointer(&iov)), 0, 0)
		if errno != 0 {
			err = errno
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return &elf.Note{
		Name:        "LINUX",
		Description: xstate[:iov.Len],
		Type:        C.NT_X86_XSTATE,
	}, nil
}
