package main

import (
	"fmt"
	"os"
	"time"

	bppir "github.com/Nv7-Github/bpp/old/ir"
	"github.com/Nv7-Github/bpp/old/parser"
)

// ParseProg parses a Go program
func ParseProg(isTiming bool, filenames []string) *parser.Program {
	files := make(map[string]string)

	// Read files
	var err error
	for _, filename := range filenames {
		dir, err := os.ReadDir(filename)
		if err == nil {
			for _, file := range dir {
				if !file.IsDir() {
					src, err := os.ReadFile(file.Name())
					handle(err)
					files[file.Name()] = string(src)
				}
			}
		} else {
			src, err := os.ReadFile(filename)
			handle(err)
			files[filename] = string(src)
		}
	}

	var out *parser.Program

	var start time.Time
	if isTiming {
		start = time.Now()
	}
	if len(files) == 1 {
		out, err = parser.Parse(filenames[0], files[filenames[0]])
		handle(err)
	} else {
		out, err = parser.ParseFiles("main.bpp", files)
		handle(err)
	}
	if isTiming {
		fmt.Println("Parsed program in", time.Since(start))
	}

	return out
}

func BuildIR(timing bool, prog *parser.Program) *bppir.IR {
	var start time.Time
	if timing {
		start = time.Now()
	}
	ir, err := bppir.CreateIR(prog)
	handle(err)
	if timing {
		fmt.Println("Built IR in", time.Since(start))
	}
	return ir
}
