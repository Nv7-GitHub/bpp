package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Nv7-Github/bpp/ir"
)

func processTools(args Args) {
	switch {
	case args.Tool.IR != nil:
		var start time.Time
		if args.Time {
			start = time.Now()
		}

		prog := ParseProg(args.Time, args.Tool.IR.Files)
		ir_v, err := ir.CreateIR(prog)
		handle(err)

		if args.Tool.IR.Output == "" {
			args.Tool.IR.Output = "main.ir"
		}

		out, err := os.Create(args.Tool.IR.Output)
		handle(err)
		defer out.Close()

		if !args.Tool.IR.Text {
			handle(ir_v.Save(out))
		} else {
			_, err := out.WriteString(ir_v.String())
			handle(err)
		}

		if args.Time {
			fmt.Println("Compiled to IR in", time.Since(start))
		}
	}
}
