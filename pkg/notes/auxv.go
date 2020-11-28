package notes

// #include <elf.h>
import "C"

import (
	"github.com/jim-minter/gcore/pkg/elf"
	"github.com/jim-minter/gcore/pkg/proc"
)

func ReadAuxv(pid int) (*elf.Note, error) {
	auxv, err := proc.ReadAuxv(pid)
	if err != nil {
		return nil, err
	}

	return &elf.Note{
		Name:        "CORE",
		Description: auxv,
		Type:        C.NT_AUXV,
	}, nil
}
