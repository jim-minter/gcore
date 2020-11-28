package proc

import (
	"fmt"
	"path/filepath"
	"strconv"
)

func Tasks(pid int) ([]int, error) {
	matches, err := filepath.Glob(fmt.Sprintf("/proc/%d/task/*", pid))
	if err != nil {
		return nil, err
	}

	tids := make([]int, 0, len(matches))
	for _, m := range matches {
		tid, err := strconv.Atoi(filepath.Base(m))
		if err != nil {
			return nil, err
		}

		tids = append(tids, tid)
	}

	return tids, nil
}
