package notes

// #include <elf.h>
import "C"

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/jim-minter/gcore/pkg/elf"
	"github.com/jim-minter/gcore/pkg/proc"
)

func File(pid int) (*elf.Note, error) {
	smaps, err := proc.ReadSmaps(pid)
	if err != nil {
		return nil, err
	}

	file := &elf.File{
		PageSize: 0x1000,
	}
	var elements []*elf.FileElement
	paths := &bytes.Buffer{}

	for _, smap := range smaps {
		if smap.Pathname == "" || smap.Pathname[0] == '[' {
			continue
		}

		elements = append(elements, &elf.FileElement{
			Start:   smap.Start,
			End:     smap.End,
			FileOfs: uint64(smap.Offset) / file.PageSize,
		})

		fmt.Fprintf(paths, "%s\x00", smap.Pathname)
	}
	file.Count = uint64(len(elements))

	buf := &bytes.Buffer{}

	err = binary.Write(buf, binary.LittleEndian, file)
	if err != nil {
		return nil, err
	}

	for _, element := range elements {
		err = binary.Write(buf, binary.LittleEndian, element)
		if err != nil {
			return nil, err
		}
	}

	_, err = paths.WriteTo(buf)
	if err != nil {
		return nil, err
	}

	return &elf.Note{
		Name:        "CORE",
		Description: buf.Bytes(),
		Type:        C.NT_FILE,
	}, nil
}
