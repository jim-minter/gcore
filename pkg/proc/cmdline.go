package proc

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func ReadCmdline(pid int) ([]byte, error) {
	b, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
	if err != nil {
		return nil, err
	}

	return bytes.ReplaceAll(bytes.TrimSpace(b), []byte{0}, []byte{' '}), nil
}
