package proc

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestReadStat(t *testing.T) {
	want := &Stat{
		Pid:        2694312,
		Comm:       "cat",
		State:      'S',
		Ppid:       2694280,
		Pgrp:       2694312,
		Session:    2694280,
		TtyNr:      34821,
		Tpgid:      2694312,
		Flags:      4194304,
		Minflt:     153,
		Priority:   20,
		NumThreads: 1,
		Starttime:  83135205,
		Vsize:      5566464,
		Rss:        987,
		Rsslim:     18446744073709551615,
		Startcode:  94015358701568,
		Endcode:    94015358716401,
		Startstack: 140722599823808,
		ExitSignal: 17,
		Processor:  6,
		StartData:  94015358728848,
		EndData:    94015358730344,
		StartBrk:   94015359299584,
		ArgStart:   140722599830419,
		ArgEnd:     140722599830423,
		EnvStart:   140722599830423,
		EnvEnd:     140722599833579,
		ExitCode:   32773,
	}

	b, err := ioutil.ReadFile("testdata/stat")
	if err != nil {
		t.Fatal(err)
	}

	got, err := readStat(b)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Error(deep.Equal(got, want))
	}
}
