package proc

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type Stat struct {
	Pid                 int32
	Comm                string
	State               byte
	Ppid                int32
	Pgrp                int32
	Session             int32
	TtyNr               int32
	Tpgid               int32
	Flags               uint32
	Minflt              uint64
	Cminflt             uint64
	Majflt              uint64
	Cmajflt             uint64
	Utime               uint64
	Stime               uint64
	Cutime              int64
	Cstime              int64
	Priority            int64
	Nice                int64
	NumThreads          int64
	Itrealvalue         int64
	Starttime           uint64
	Vsize               uint64
	Rss                 int64
	Rsslim              uint64
	Startcode           uint64
	Endcode             uint64
	Startstack          uint64
	Kstkesp             uint64
	Kstkeip             uint64
	Signal              uint64
	Blocked             uint64
	Sigignore           uint64
	Sigcatch            uint64
	Wchan               uint64
	Nswap               uint64
	Cnswap              uint64
	ExitSignal          int32
	Processor           int32
	RtPriority          uint32
	Policy              uint32
	DelayacctBlkioTicks uint64
	GuestTime           uint64
	CguestTime          int64
	StartData           uint64
	EndData             uint64
	StartBrk            uint64
	ArgStart            uint64
	ArgEnd              uint64
	EnvStart            uint64
	EnvEnd              uint64
	ExitCode            int32
}

func ReadStat(pid, tid int) (*Stat, error) {
	path := fmt.Sprintf("/proc/%d/task/%d/stat", pid, tid)
	if tid == 0 {
		path = fmt.Sprintf("/proc/%d/stat", pid)
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return readStat(b)
}

func readStat(b []byte) (*Stat, error) {
	values := strings.Fields(string(b))

	stat := &Stat{}

	v := reflect.ValueOf(stat).Elem()

	for i := 0; i < v.NumField() && i < len(values); i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.Int32, reflect.Int64:
			v, err := strconv.ParseInt(values[i], 10, 64)
			if err != nil {
				return nil, err
			}
			f.SetInt(v)

		case reflect.Uint32, reflect.Uint64:
			v, err := strconv.ParseUint(values[i], 10, 64)
			if err != nil {
				return nil, err
			}
			f.SetUint(v)

		case reflect.String:
			f.SetString(values[i])

		case reflect.Uint8:
			f.SetUint(uint64(values[i][0]))

		default:
			panic(1)
		}
	}

	stat.Comm = stat.Comm[1 : len(stat.Comm)-1]

	return stat, nil
}
