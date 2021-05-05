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
	cppname := name + ".cpp"
	err = os.WriteFile(cppname, []byte(compiled), os.ModePerm)
	handle(err)

	if args.Time {
		start = time.Now()
	}
	outFile := args.Build.Output
	if outFile == "" {
		outFile = name
		if runtime.GOOS == "windows" {
			outFile += ".exe"
		}
	}
	cmd := exec.Command(args.Build.CC, "-lstdc++", "-o", outFile, "-O3", cppname)

	stderr := bytes.NewBuffer(make([]byte, 0))
	cmd.Stderr = stderr

	err = cmd.Run()
	if err != nil {
		handle(errors.New(stderr.String()))
	}
	if args.Time {
		fmt.Println("Compiled in", time.Since(start))
	}

	if !args.Build.Preserve {
		err = os.Remove(cppname)
		handle(err)
	}
}
