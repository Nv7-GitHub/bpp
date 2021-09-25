package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/Nv7-Github/bpp/old/compiler"
	"github.com/Nv7-Github/bpp/old/membuild"
	"github.com/Nv7-Github/bpp/parser"
)

// CompileCmd implements the "compile" sub-command
func CompileCmd(args Args, prog *parser.Program) {
	ext := filepath.Ext(args.Old.Build.Files[0])
	name := args.Old.Build.Files[0][0 : len(args.Old.Build.Files[0])-len(ext)]

	var start time.Time
	if args.Time {
		start = time.Now()
	}
	compiled, err := compiler.Compile(prog)
	handle(err)
	if args.Time {
		fmt.Println("Built in", time.Since(start))
	}
	llname := name + ".ll"
	err = os.WriteFile(llname, []byte(compiled), os.ModePerm)
	handle(err)

	if args.Time {
		start = time.Now()
	}

	if args.Old.Build.CC != "" {
		outFile := args.Old.Build.Output
		if outFile == "" {
			outFile = name
			if runtime.GOOS == "windows" {
				outFile += ".exe"
			}
		}
		cmd := exec.Command(args.Old.Build.CC, "-o", outFile, "-O3", llname)

		stderr := bytes.NewBuffer(make([]byte, 0))
		cmd.Stderr = stderr

		err = cmd.Run()
		if err != nil {
			handle(errors.New(stderr.String()))
		}
		if args.Time {
			fmt.Println("Compiled in", time.Since(start))
		}
	} else if args.Old.Build.Asm != "" {
		outFile := args.Old.Build.Output
		if outFile == "" {
			outFile = name
		}
		outFile += ".s"

		cmd := exec.Command(args.Old.Build.Asm, "-o", outFile, "-O3", llname)

		stderr := bytes.NewBuffer(make([]byte, 0))
		cmd.Stderr = stderr

		err = cmd.Run()
		if err != nil {
			handle(errors.New(stderr.String()))
		}
		if args.Time {
			fmt.Println("Created assembly in", time.Since(start))
		}
	}

	if args.Old.Build.CC != "" || args.Old.Build.Asm != "" {
		err = os.Remove(llname)
		handle(err)
	}
}

// MembuildCmd implements the "membuild" sub-command
func MembuildCmd(args Args, prog *parser.Program) {
	var start time.Time
	if args.Time {
		start = time.Now()
	}
	p, err := membuild.Build(prog)
	handle(err)
	if args.Time {
		fmt.Println("Built in", time.Since(start))
	}

	if len(args.Old.Membuild.Args) > 0 {
		p.Args = strings.Split(args.Old.Membuild.Args, ",")
	} else {
		p.Args = make([]string, 0)
	}

	if args.Time {
		start = time.Now()
	}

	p.Runner = func(d membuild.Data) error {
		if !d.Type.IsEqual(parser.NULL) && d.Value != "" {
			_, err := fmt.Println(d.Value)
			return err
		}
		return nil
	}

	err = p.Run()
	handle(err)
	if args.Time {
		fmt.Println("Ran in", time.Since(start))
	}
}

func OldCmd(args Args) {
	switch {
	case args.Old.Build != nil:
		fmt.Println("WARNING: Build is a legacy command and may be removed at any time!")

		if len(args.Old.Build.Files) < 1 {
			handle(errors.New("you must supply at least one file"))
		}
		prog := ParseProg(args.Time, args.Old.Build.Files)
		CompileCmd(args, prog)

	case args.Old.Membuild != nil:
		fmt.Println("WARNING: Membuild is a legacy command and may be removed at any time!")

		if len(args.Old.Membuild.Files) < 1 {
			handle(errors.New("you must supply at least one file"))
		}
		prog := ParseProg(args.Time, args.Old.Membuild.Files)
		MembuildCmd(args, prog)
	}
}
