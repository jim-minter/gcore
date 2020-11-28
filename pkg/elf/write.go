package elf

import (
	"bytes"
	"debug/elf"
	"encoding/binary"
	"io"
)

const align = 0x1000

type ident struct {
	Magic      [4]byte
	Class      elf.Class
	Data       elf.Data
	Version    elf.Version
	OSABI      elf.OSABI
	ABIVersion uint8
	_          [7]byte
}

func newHeader(f *elf.File) (*elf.Header64, error) {
	i := ident{
		Class:   elf.ELFCLASS64,
		Data:    elf.ELFDATA2LSB,
		Version: elf.EV_CURRENT,
	}
	copy(i.Magic[:], []byte(elf.ELFMAG))

	h := &elf.Header64{
		Type:    uint16(f.Type),
		Machine: uint16(elf.EM_X86_64),
		Version: uint32(i.Version),
		Ehsize:  uint16(binary.Size(&elf.Header64{})),
	}

	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, i)
	if err != nil {
		return nil, err
	}
	copy(h.Ident[:], buf.Bytes())

	calcProgs(f, h)

	return h, nil
}

func calcProgs(f *elf.File, h *elf.Header64) {
	if len(f.Progs) == 0 {
		return
	}

	h.Phoff = uint64(h.Ehsize)
	h.Phentsize = uint16(binary.Size(&elf.Prog64{}))
	h.Phnum = uint16(len(f.Progs))

	base := int(h.Ehsize)
	base += len(f.Progs) * int(h.Phentsize)

	for i := range f.Progs {
		f.Progs[i].Off = uint64(base)
		base += int(f.Progs[i].Filesz)

		base = (base + align - 1) & ^(align - 1)
	}
}

func Write(w io.Writer, f *elf.File) error {
	h, err := newHeader(f)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.LittleEndian, h)
	if err != nil {
		return err
	}

	off := uint64(binary.Size(h))

	for i, prog := range f.Progs {
		ph := &elf.Prog64{
			Type:   uint32(prog.Type),
			Flags:  uint32(prog.Flags),
			Off:    prog.Off,
			Vaddr:  prog.Vaddr,
			Filesz: prog.Filesz,
			Memsz:  prog.Memsz,
		}

		if i > 0 {
			ph.Align = align
		}

		err = binary.Write(w, binary.LittleEndian, ph)
		if err != nil {
			return err
		}

		off += uint64(binary.Size(ph))
	}

	for _, prog := range f.Progs {
		if prog.Filesz == 0 {
			continue
		}

		for prog.Off > off {
			n, err := w.Write(make([]byte, min(prog.Off-off, align)))
			if err != nil {
				return err
			}

			off += uint64(n)
		}

		n, err := io.Copy(w, prog.ReaderAt.(io.Reader))
		if err != nil {
			return err
		}

		off += uint64(n)
	}

	return nil
}

func min(i, j uint64) uint64 {
	if i < j {
		return i
	}

	return j
}
