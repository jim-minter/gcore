package proc

import (
	"fmt"
	"io/ioutil"
)

func ReadAuxv(pid int) ([]byte, error) {
	return ioutil.ReadFile(fmt.Sprintf("/proc/%d/auxv", pid))
}
