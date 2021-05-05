package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Nv7-Github/Bpp/parser"
	arg "github.com/alexflint/go-arg"
)

var p *arg.Parser

func handle(err error) {
	if err != nil {
		p.Fail(err.Error())
	}
}

type Build struct {
	Output   string `help:"output file for executable" arg:"-o"`
	CC       string `default:"cc" help:"C++ compiler"`
	Preserve bool   `help:"preserve C++ source code" arg:"-p"`
	File     string `arg:"positional,-i,--input" help:"input B++ program"`
}

type Run struct {
	Args []string `help:"arguments for program"`
	File string   `arg:"positional,-i,--input" help:"input B++ program"`
}

type Args struct {
	Build *Build `arg:"subcommand:build"`
	Run   *Run   `arg:"subcommand:run"`
	Time  bool   `help:"print timing for each stage" arg:"-t"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var args Args
	p = arg.MustParse(&args)

	switch {
	case args.Build != nil:
		prog := ParseProg(args.Time, args.Build.File)
		CompileCmd(args, prog)
	case args.Run != nil:
		prog := ParseProg(args.Time, args.Run.File)
		RunCmd(args, prog)
	default:
		p.WriteUsage(os.Stdout)
	}
}

func ParseProg(isTiming bool, filename string) *parser.Program {
	src, err := os.ReadFile(filename)
	handle(err)

	var start time.Time
	if isTiming {
		start = time.Now()
	}
	prog, err := parser.Parse(string(src))
	handle(err)
	if isTiming {
		fmt.Println("Parsed program in", time.Since(start))
	}

	return prog
}
