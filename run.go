package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Nv7-Github/Bpp/membuild"
	"github.com/Nv7-Github/Bpp/parser"
)

func RunCmd(args Args, prog *parser.Program) {
	var start time.Time
	if args.Time {
		start = time.Now()
	}
	p, err := membuild.Build(prog)
	handle(err)
	if args.Time {
		fmt.Println("Built in", time.Since(start))
	}

	if args.Time {
		start = time.Now()
	}
	err = membuild.RunProgram(p, os.Stdout)
	handle(err)
	if args.Time {
		fmt.Println("Ran in", time.Since(start))
	}
}
