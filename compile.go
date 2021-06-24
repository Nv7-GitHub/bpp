package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/Nv7-Github/Bpp/compiler"
	"github.com/Nv7-Github/Bpp/parser"
)

// CompileCmd implements the "compile" sub-command
func CompileCmd(args Args, prog *parser.Program) {
	ext := filepath.Ext(args.Build.File)
	name := args.Build.File[0 : len(args.Build.File)-len(ext)]

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

	if args.Build.CC != "" {
		outFile := args.Build.Output
		if outFile == "" {
			outFile = name
			if runtime.GOOS == "windows" {
				outFile += ".exe"
			}
		}
		cmd := exec.Command(args.Build.CC, "-o", outFile, "-O3", llname)

		stderr := bytes.NewBuffer(make([]byte, 0))
		cmd.Stderr = stderr

		err = cmd.Run()
		if err != nil {
			handle(errors.New(stderr.String()))
		}
		if args.Time {
			fmt.Println("Compiled in", time.Since(start))
		}
	} else if args.Build.Asm != "" {
		outFile := args.Build.Output
		if outFile == "" {
			outFile = name
		}
		outFile += ".s"

		cmd := exec.Command(args.Build.Asm, "-o", outFile, "-O3", llname)

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

	if args.Build.CC != "" || args.Build.Asm != "" {
		err = os.Remove(llname)
		handle(err)
	}
}
