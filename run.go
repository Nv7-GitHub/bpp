package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Nv7-Github/bpp/membuild"
	"github.com/Nv7-Github/bpp/parser"
)

// RunCmd implements the "run" sub-command
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

	if len(args.Run.Args) > 0 {
		p.Args = strings.Split(args.Run.Args, ",")
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
