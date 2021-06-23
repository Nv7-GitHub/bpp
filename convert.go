package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"time"

	"github.com/Nv7-Github/Bpp/gobpp"
)

func ConvertCmd(args Args) {
	src, err := os.ReadFile(args.Convert.File)
	handle(err)

	ext := filepath.Ext(args.Convert.File)
	name := args.Convert.File[0 : len(args.Convert.File)-len(ext)]

	var start time.Time
	if args.Time {
		start = time.Now()
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	if args.Time {
		fmt.Println("Parsed in", time.Since(start))
		start = time.Now()
	}

	out, err := gobpp.Convert(fset, args.Convert.File, f)
	handle(err)

	if args.Time {
		fmt.Println("Converted in", time.Since(start))
	}

	outFile := args.Convert.Output
	if outFile == "" {
		outFile = name + ".bpp"
	}
	err = os.WriteFile(outFile, []byte(out), os.ModePerm)
	handle(err)
}
