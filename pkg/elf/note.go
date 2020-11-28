package elf

import (
	"encoding/binary"
	"io"
)

type Note struct {
	Name        string
	Description []byte
	Type        uint32
}

func (n *Note) Write(w io.Writer) error {
	name := n.Name + "\x00"

	err := binary.Write(w, binary.LittleEndian, struct {
		Namesz uint32
		Descsz uint32
		Type   uint32
	}{
		Namesz: uint32(len(name)),
		Descsz: uint32(len(n.Description)),
		Type:   n.Type,
	})
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(name))
	if err != nil {
		return err
	}

	_, err = w.Write(make([]byte, (4-len(name)&3)&3))
	if err != nil {
		return err
	}

	_, err = w.Write(n.Description)
	if err != nil {
		return err
	}

	_, err = w.Write(make([]byte, (4-len(n.Description)&3)&3))
	return err
}
