package gcore

import (
	"bytes"
	"debug/elf"
	"io"
	"os"

	pkgelf "github.com/jim-minter/gcore/pkg/elf"
	pkgnotes "github.com/jim-minter/gcore/pkg/notes"
	"github.com/jim-minter/gcore/pkg/proc"
	"github.com/jim-minter/gcore/pkg/ptrace"
)

func notes(pid int, tids []int) (*elf.Prog, error) {
	buf := &bytes.Buffer{}

	n, err := pkgnotes.Prpsinfo(pid)
	if err != nil {
		return nil, err
	}

	err = n.Write(buf)
	if err != nil {
		return nil, err
	}

	for _, tid := range tids {
		for _, f := range []func(int, int) (*pkgelf.Note, error){
			pkgnotes.Prstatus,
			pkgnotes.Fpregset,
			pkgnotes.Xstate,
			pkgnotes.Siginfo,
		} {
			n, err = f(pid, tid)
			if err != nil {
				return nil, err
			}

			err = n.Write(buf)
			if err != nil {
				return nil, err
			}
		}
	}

	for _, f := range []func(int) (*pkgelf.Note, error){
		pkgnotes.ReadAuxv,
		pkgnotes.File,
	} {
		n, err = f(pid)
		if err != nil {
			return nil, err
		}

		err = n.Write(buf)
		if err != nil {
			return nil, err
		}
	}

	return &elf.Prog{
		ProgHeader: elf.ProgHeader{
			Type:   elf.PT_NOTE,
			Filesz: uint64(buf.Len()),
		},
		ReaderAt: bytes.NewReader(buf.Bytes()),
	}, nil
}

func progs(pid int) (progs []*elf.Prog, err error) {
	mem, err := proc.Mem(pid)
	if err != nil {
		return nil, err
	}

	smaps, err := proc.ReadSmaps(pid)
	if err != nil {
		return nil, err
	}

	for _, smap := range smaps {
		prog := &elf.Prog{
			ProgHeader: elf.ProgHeader{
				Type:  elf.PT_LOAD,
				Vaddr: smap.Start,
				Memsz: smap.End - smap.Start,
			},
		}

		if smap.Perms&proc.PermR != 0 {
			prog.Flags |= elf.PF_R
		}
		if smap.Perms&proc.PermW != 0 {
			prog.Flags |= elf.PF_W
		}
		if smap.Perms&proc.PermX != 0 {
			prog.Flags |= elf.PF_X
		}

		if !smap.HasVMFlag("dd") &&
			!smap.HasVMFlag("io") &&
			smap.Perms&proc.PermR != 0 &&
			int64(smap.Start) >= 0 /* TODO: hack */ {
			prog.Filesz = prog.Memsz
			prog.ReaderAt = io.NewSectionReader(mem, int64(smap.Start), int64(smap.End-smap.Start))
		}

		progs = append(progs, prog)
	}

	return progs, nil
}

func Run(pid int) error {
	tids, err := ptrace.Seize(pid)
	if err != nil {
		return err
	}

	notes, err := notes(pid, tids)
	if err != nil {
		return err
	}

	progs, err := progs(pid)
	if err != nil {
		return err
	}

	f := &elf.File{
		FileHeader: elf.FileHeader{
			Type: elf.ET_CORE,
		},
		Progs: append([]*elf.Prog{notes}, progs...),
	}

	return pkgelf.Write(os.Stdout, f)
}
