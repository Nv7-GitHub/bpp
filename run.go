package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Nv7-Github/bpp/ir"
	"github.com/Nv7-Github/bpp/parser"
	"github.com/Nv7-Github/bpp/run"
)

// RunCmd implements the "run" sub-command
func RunCmd(args Args, prog *parser.Program) {
	var start time.Time
	if args.Time {
		start = time.Now()
	}
	ir, err := ir.CreateIR(prog)
	handle(err)
	runnable := run.NewRunnable(ir)
	runnable.Stdout = os.Stdout
	if args.Time {
		fmt.Println("Built IR in", time.Since(start))
	}

	var progArgs []string
	if len(args.Run.Args) > 0 {
		progArgs = strings.Split(args.Run.Args, ",")
	} else {
		progArgs = make([]string, 0)
	}

	if args.Time {
		start = time.Now()
	}

	err = runnable.Run(progArgs)
	handle(err)
	if args.Time {
		fmt.Println("Ran in", time.Since(start))
	}
}
