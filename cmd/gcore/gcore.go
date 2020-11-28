package main

// extern int pid;
import "C"

import (
	"fmt"
	"os"

	"github.com/jim-minter/gcore/pkg/gcore"
)

func main() {
	if err := gcore.Run(int(C.pid)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
