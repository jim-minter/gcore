package notes

// #include <sys/procfs.h>
// #include <elf.h>
import "C"

import (
	"strings"
	"syscall"
	"unsafe"

	"github.com/jim-minter/gcore/pkg/elf"
	"github.com/jim-minter/gcore/pkg/proc"
)

func Prpsinfo(pid int) (*elf.Note, error) {
	stat, err := proc.ReadStat(pid, 0)
	if err != nil {
		return nil, err
	}

	fi, err := proc.ProcStat(pid)
	if err != nil {
		return nil, err
	}

	cmdline, err := proc.ReadCmdline(pid)
	if err != nil {
		return nil, err
	}

	prpsinfo := &C.struct_elf_prpsinfo{
		pr_state: C.char(strings.IndexByte("RSDTtZXxKWP", stat.State)),
		pr_sname: C.char(stat.State),
		pr_nice:  C.char(stat.Nice),
		pr_flag:  C.ulong(stat.Flags),
		pr_uid:   C.uint(fi.Sys().(*syscall.Stat_t).Uid),
		pr_gid:   C.uint(fi.Sys().(*syscall.Stat_t).Gid),
		pr_pid:   C.int(stat.Pid),
		pr_ppid:  C.int(stat.Ppid),
		pr_pgrp:  C.int(stat.Pgrp),
		pr_sid:   C.int(stat.Session),
	}

	copy((*(*[unsafe.Sizeof(prpsinfo.pr_fname)]byte)(unsafe.Pointer(&prpsinfo.pr_fname)))[:], []byte(stat.Comm))
	copy((*(*[unsafe.Sizeof(prpsinfo.pr_psargs)]byte)(unsafe.Pointer(&prpsinfo.pr_psargs)))[:], cmdline)

	if stat.State == 'Z' {
		prpsinfo.pr_zomb = 1
	}

	return &elf.Note{
		Name:        "CORE",
		Description: C.GoBytes(unsafe.Pointer(prpsinfo), C.int(unsafe.Sizeof(*prpsinfo))),
		Type:        C.NT_PRPSINFO,
	}, nil
}
