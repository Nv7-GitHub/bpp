package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/Nv7-Github/bpp/builder"
	"github.com/Nv7-Github/bpp/parser"
	"github.com/Nv7-Github/bpp/run"
)

// RunCmd implements the "run" sub-command
func RunCmd(args Args, prog *parser.Program) {
	var start time.Time
	ir := BuildIR(args.Time, prog)

	var progArgs []string
	if len(args.Run.Args) > 0 {
		progArgs = strings.Split(args.Run.Args, ",")
	} else {
		progArgs = make([]string, 0)
	}

	if args.Time {
		start = time.Now()
	}

	runnable := run.NewRunnable(ir)
	runnable.Stdout = os.Stdout
	err := runnable.Run(progArgs)
	handle(err)

	if args.Time {
		fmt.Println("Ran in", time.Since(start))
	}
}

func BuildCmd(args Args) {
	if len(args.Build.Files) < 1 {
		handle(errors.New("you must supply at least one file"))
	}
	prog := ParseProg(args.Time, args.Build.Files)
	ir := BuildIR(args.Time, prog)

	var start time.Time
	if args.Time {
		start = time.Now()
	}

	llvm, err := builder.Build(ir)
	handle(err)

	if args.Time {
		fmt.Println("Built LLVM in", time.Since(start))
	}

	if args.Time {
		start = time.Now()
	}

	if args.Build.LLVM {
		if args.Build.Output == "" {
			args.Build.Output = "main.ll"
		}

		out, err := os.Create(args.Build.Output)
		handle(err)
		defer out.Close()

		_, err = out.WriteString(llvm)
		handle(err)
	} else {
		tmpFile, err := os.CreateTemp("", "")
		handle(err)
		defer tmpFile.Close()

		_, err = tmpFile.WriteString(llvm)
		handle(err)

		if args.Build.Output == "" {
			args.Build.Output = "a.out"
		}

		// Build
		cmd := exec.Command("clang", tmpFile.Name(), "-o", args.Build.Output)
		err = cmd.Run()
		handle(err)
	}

	if args.Time {
		fmt.Println("Created output in", time.Since(start))
	}
}
