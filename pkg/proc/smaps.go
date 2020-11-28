package proc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Smap struct {
	Start    uint64
	End      uint64
	Perms    Perm
	Offset   uint32
	Dev      string
	Inode    uint32
	Pathname string

	Data map[string]string
}

func (smap *Smap) HasVMFlag(flag string) bool {
	vmflags, ok := smap.Data["vmflags"]
	if !ok {
		return false
	}

	flag = strings.ToLower(flag)

	flags := strings.Split(vmflags, " ")
	for _, f := range flags {
		if f == flag {
			return true
		}
	}

	return false
}

type Perm int

const (
	PermR Perm = 1 << iota
	PermW
	PermX
	PermS
	PermP
)

var header = regexp.MustCompile(`^([0-9a-f]{0,16})-([0-9a-f]{0,16}) ([-r][-w][-x][-sp]) ([0-9a-f]{8}) ([0-9a-f]{2}:[0-9a-f]{2}) ([0-9]+) +([^ ]+)?`)

func ReadSmaps(pid int) ([]*Smap, error) {
	f, err := os.Open(fmt.Sprintf("/proc/%d/smaps", pid))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readSmaps(f)
}

func readSmaps(r io.Reader) (smaps []*Smap, err error) {
	var smap *Smap

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		m := header.FindStringSubmatch(scanner.Text())
		if m != nil {
			if smap != nil {
				smaps = append(smaps, smap)
				smap = nil
			}

			start, err := strconv.ParseUint(m[1], 16, 64)
			if err != nil {
				return nil, err
			}

			end, err := strconv.ParseUint(m[2], 16, 64)
			if err != nil {
				return nil, err
			}

			offset, err := strconv.ParseUint(m[4], 16, 32)
			if err != nil {
				return nil, err
			}

			inode, err := strconv.ParseUint(m[6], 10, 32)
			if err != nil {
				return nil, err
			}

			smap = &Smap{
				Start:    start,
				End:      end,
				Offset:   uint32(offset),
				Dev:      m[5],
				Inode:    uint32(inode),
				Pathname: m[7],

				Data: map[string]string{},
			}

			if strings.ContainsRune(m[3], 'r') {
				smap.Perms |= PermR
			}
			if strings.ContainsRune(m[3], 'w') {
				smap.Perms |= PermW
			}
			if strings.ContainsRune(m[3], 'x') {
				smap.Perms |= PermX
			}
			if strings.ContainsRune(m[3], 'p') {
				smap.Perms |= PermP
			}
			if strings.ContainsRune(m[3], 's') {
				smap.Perms |= PermS
			}

			continue
		}

		kv := strings.SplitN(scanner.Text(), ":", 2)

		k := strings.ToLower(strings.TrimSpace(kv[0]))
		v := strings.ToLower(strings.TrimSpace(kv[1]))

		if _, ok := smap.Data[k]; ok {
			return nil, fmt.Errorf("duplicate data key")
		}

		smap.Data[k] = v
	}
	if scanner.Err() != nil {
		return nil, err
	}

	if smap != nil {
		smaps = append(smaps, smap)
		smap = nil
	}

	return smaps, nil
}
