package ptrace

import (
	"golang.org/x/sys/unix"

	"github.com/jim-minter/gcore/pkg/proc"
)

func Seize(pid int) ([]int, error) {
	seized := map[int]struct{}{}

	for {
		var didWork bool

		tids, err := proc.Tasks(pid)
		if err != nil {
			return nil, err
		}

		for _, tid := range tids {
			if _, ok := seized[tid]; ok {
				continue
			}

			err = Do(func() error { return unix.PtraceSeize(tid) })
			if err != nil {
				return nil, err
			}

			err = Do(func() error { return unix.PtraceInterrupt(tid) })
			if err != nil {
				return nil, err
			}

			seized[tid] = struct{}{}
			didWork = true
		}

		if !didWork {
			break
		}
	}

	tids := make([]int, 0, len(seized))
	for tid := range seized {
		tids = append(tids, tid)
	}

	return tids, nil
}
