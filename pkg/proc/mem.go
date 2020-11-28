package proc

import (
	"fmt"
	"os"
)

func Mem(pid int) (*os.File, error) {
	return os.Open(fmt.Sprintf("/proc/%d/mem", pid))
}
