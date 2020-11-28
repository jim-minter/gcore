package proc

import (
	"fmt"
	"os"
)

func ProcStat(pid int) (os.FileInfo, error) {
	return os.Stat(fmt.Sprintf("/proc/%d", pid))
}
