package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	"github.com/Nv7-Github/Bpp/ir"
	"github.com/Nv7-Github/Bpp/parser"
	arg "github.com/alexflint/go-arg"
)

var p *arg.Parser

func handle(err error) {
	if err != nil {
		p.Fail(err.Error())
	}
}

// Build defines the "build" sub-command
type Build struct {
	Output string   `help:"output file for executable" arg:"-o"`
	CC     string   `help:"LLVM compiler (optional)"`
	Asm    string   `help:"LLC compiler (optional)"`
	Files  []string `arg:"positional,-i,--input" help:"input B++ program"`
}

// Run defines the "run" sub-command
type Run struct {
	Args  string   `help:"arguments for program, comma-seperated"`
	Files []string `arg:"positional,-i,--input" help:"input B++ program"`
}

// Convert defines the "convert" sub-command
type Convert struct {
	Output string `help:"output B++ program" arg:"-o"`
	File   string `arg:"positional,-i,--input" help:"input Go program"`
}

// IR defines the "ir" sub-command
type IR struct {
	Files []string `arg:"positional,-i,--input" help:"input B++ program"`
}

// Args defines the program's arguments
type Args struct {
	Build   *Build   `arg:"subcommand:build" help:"compile a B++ program"`
	Run     *Run     `arg:"subcommand:run" help:"run a B++ program"`
	IR      *IR      `arg:"subcommand:ir" help:"generate B++ IR"`
	Convert *Convert `arg:"subcommand:convert" help:"convert a go program to a B++ program"`
	Time    bool     `help:"print timing for each stage" arg:"-t"`

	CPUProf string `help:"CPU pprof statistics output file"`
	Memprof string `help:"heap pprof statistics output file"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var args Args
	p = arg.MustParse(&args)

	if args.CPUProf != "" {
		pprofFile, err := os.OpenFile(args.CPUProf, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		handle(err)
		err = pprof.StartCPUProfile(pprofFile)
		handle(err)
		defer pprof.StopCPUProfile()
	}

	switch {
	case args.Build != nil:
		if len(args.Build.Files) < 1 {
			handle(errors.New("you must supply at least one file"))
		}
		prog := ParseProg(args.Time, args.Build.Files)
		CompileCmd(args, prog)
	case args.Run != nil:
		if len(args.Run.Files) < 1 {
			handle(errors.New("you must supply at least one file"))
		}
		prog := ParseProg(args.Time, args.Run.Files)
		RunCmd(args, prog)
	case args.Convert != nil:
		ConvertCmd(args)
	case args.IR != nil:
		prog := ParseProg(args.Time, args.IR.Files)
		ir, err := ir.CreateIR(prog)
		handle(err)
		fmt.Println(ir.String())
	default:
		p.WriteUsage(os.Stdout)
	}

	if args.Memprof != "" {
		pprofFile, err := os.OpenFile(args.Memprof, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		handle(err)
		err = pprof.WriteHeapProfile(pprofFile)
		handle(err)
	}
}

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
