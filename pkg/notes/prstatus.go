package notes

// #include <sys/procfs.h>
// #include <sys/ptrace.h>
// #include <elf.h>
import "C"

import (
	"syscall"
	"unsafe"

	elf "github.com/jim-minter/gcore/pkg/elf"
	"github.com/jim-minter/gcore/pkg/proc"
	"github.com/jim-minter/gcore/pkg/ptrace"
)

func Prstatus(pid, tid int) (*elf.Note, error) {
	stat, err := proc.ReadStat(pid, tid)
	if err != nil {
		return nil, err
	}

	prstatus := &C.struct_elf_prstatus{
		pr_sigpend: C.ulong(stat.Signal),
		pr_sighold: C.ulong(stat.Blocked),
		pr_pid:     C.int(stat.Pid),
		pr_ppid:    C.int(stat.Ppid),
		pr_pgrp:    C.int(stat.Pgrp),
		pr_sid:     C.int(stat.Session),
		pr_utime:   C.struct_timeval{tv_sec: C.long(stat.Utime) / 1000000, tv_usec: C.long(stat.Utime) % 1000000},
		pr_stime:   C.struct_timeval{tv_sec: C.long(stat.Stime) / 1000000, tv_usec: C.long(stat.Stime) % 1000000},
		pr_cutime:  C.struct_timeval{tv_sec: C.long(stat.Cutime) / 1000000, tv_usec: C.long(stat.Cutime) % 1000000},
		pr_cstime:  C.struct_timeval{tv_sec: C.long(stat.Cstime) / 1000000, tv_usec: C.long(stat.Cstime) % 1000000},
		pr_fpvalid: 1,
	}

	err = ptrace.Do(func() (err error) {
		_, _, errno := syscall.Syscall6(syscall.SYS_PTRACE, C.PTRACE_GETREGS, uintptr(tid), 0, uintptr(unsafe.Pointer(&prstatus.pr_reg)), 0, 0)
		if errno != 0 {
			err = errno
		}
		return err
	})

	return &elf.Note{
		Name:        "CORE",
		Description: C.GoBytes(unsafe.Pointer(prstatus), C.int(unsafe.Sizeof(*prstatus))),
		Type:        C.NT_PRSTATUS,
	}, nil
}
