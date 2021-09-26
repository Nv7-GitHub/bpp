package main

import (
	"errors"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	arg "github.com/alexflint/go-arg"
)

var p *arg.Parser

func handle(err error) {
	if err != nil {
		p.Fail(err.Error())
	}
}

// Build_Old defines the old "build" sub-command
type Build_Old struct {
	Output string   `help:"output file for executable" arg:"-o"`
	CC     string   `help:"LLVM compiler (optional)"`
	Asm    string   `help:"LLC compiler (optional)"`
	Files  []string `arg:"positional,-i,--input" help:"input B++ program"`
}

// Build defines the "build" sub-command
type Build struct {
	Output string   `help:"output file for executable or LLVM" arg:"-o"`
	LLVM   bool     `help:"whether to produce LLVM IR or an executable" arg:"-l"`
	Files  []string `arg:"positional,-i,--input" help:"input B++ program"`
}

// Membuild defines the "membuild" sub-command
type Membuild struct {
	Args  string   `help:"arguments for program, comma-seperated"`
	Files []string `arg:"positional,-i,--input" help:"input B++ program"`
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

// Tools defines the "tools" sub-command
type Tool struct {
	IR *IR `arg:"subcommand:ir" help:"generate B++ ir"`
}

// IR defines the "ir" sub-command
type IR struct {
	Files  []string `arg:"positional,-i,--input" help:"input B++ program"`
	Output string   `help:"output file for IR" arg:"-o"`
	Text   bool     `help:"text format"`
}

// Old defines the "old" sub-command
type Old struct {
	Membuild *Membuild  `arg:"subcommand:membuild" help:"run a B++ program using the legacy interpreter"`
	Build    *Build_Old `arg:"subcommand:build" help:"compile a B++ program using the legacy compiler"`
}

// Args defines the program's arguments
type Args struct {
	Run     *Run     `arg:"subcommand:run" help:"run a B++ program"`
	Build   *Build   `arg:"subcommand:build" help:"compile a B++ program"`
	Convert *Convert `arg:"subcommand:convert" help:"convert a go program to a B++ program"`
	Tool    *Tool    `arg:"subcommand:tool" help:"run B++ individual tools"`
	Old     *Old     `arg:"subcommand:old" help:"run a legacy B++ tool"`
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
	case args.Old != nil:
		OldCmd(args)

	case args.Run != nil:
		if len(args.Run.Files) < 1 {
			handle(errors.New("you must supply at least one file"))
		}
		prog := ParseProg(args.Time, args.Run.Files)
		RunCmd(args, prog)
	case args.Build != nil:
		BuildCmd(args)
	case args.Convert != nil:
		ConvertCmd(args)
	case args.Tool != nil:
		processTools(args)
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
