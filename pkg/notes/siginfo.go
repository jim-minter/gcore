package notes

// #include <sys/ptrace.h>
// #include <elf.h>
// #include <signal.h>
import "C"

import (
	"syscall"
	"unsafe"

	"github.com/jim-minter/gcore/pkg/elf"
	"github.com/jim-minter/gcore/pkg/ptrace"
)

func Siginfo(pid, tid int) (*elf.Note, error) {
	siginfo := &C.siginfo_t{}
	args := &C.struct___ptrace_peeksiginfo_args{
		nr: 1,
	}

	err := ptrace.Do(func() (err error) {
		_, _, errno := syscall.Syscall6(syscall.SYS_PTRACE, C.PTRACE_PEEKSIGINFO, uintptr(tid), uintptr(unsafe.Pointer(args)), uintptr(unsafe.Pointer(siginfo)), 0, 0)
		if errno != 0 {
			err = errno
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return &elf.Note{
		Name:        "CORE",
		Description: C.GoBytes(unsafe.Pointer(siginfo), C.int(unsafe.Sizeof(*siginfo))),
		Type:        C.NT_SIGINFO,
	}, nil
}
