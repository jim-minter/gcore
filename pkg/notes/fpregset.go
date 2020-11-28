package notes

// #include <sys/ptrace.h>
// #include <sys/user.h>
// #include <elf.h>
import "C"

import (
	"syscall"
	"unsafe"

	"github.com/jim-minter/gcore/pkg/elf"
	"github.com/jim-minter/gcore/pkg/ptrace"
)

func Fpregset(pid, tid int) (*elf.Note, error) {
	fpregset := &C.struct_user_fpregs_struct{}

	err := ptrace.Do(func() (err error) {
		_, _, errno := syscall.Syscall6(syscall.SYS_PTRACE, C.PTRACE_GETFPREGS, uintptr(tid), 0, uintptr(unsafe.Pointer(fpregset)), 0, 0)
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
		Description: C.GoBytes(unsafe.Pointer(fpregset), C.int(unsafe.Sizeof(*fpregset))),
		Type:        C.NT_FPREGSET,
	}, nil
}
