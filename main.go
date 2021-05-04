package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Nv7-Github/Bpp/membuild"
	"github.com/Nv7-Github/Bpp/parser"
)

var filename string

func init() {
	flag.StringVar(&filename, "file", "", "File to execute using bpp.")
}

func handle(err error) {
	if err != nil {
		fmt.Println("Error! " + err.Error())
		os.Exit(0)
	}
}

func main() {
	flag.Parse()

	if filename == "" {
		flag.Usage()
		return
	}

	script, err := ioutil.ReadFile(filename)
	handle(err)
	src := strings.TrimSpace(string(script))
	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	prog, err := parser.Parse(src)
	handle(err)
	fmt.Println("Parsed in", time.Since(start))

	start = time.Now()
	built, err := membuild.Build(prog)
	handle(err)
	fmt.Println("Built in", time.Since(start))

	built.Args = []string{"2"}

	start = time.Now()
	i := 0
	for !(i == len(built.Instructions)) {
		val, err := built.Instructions[i](built)
		handle(err)
		if val.Type == membuild.GOTO {
			i = val.Value.(int)
			continue
		}
		if val.Type != parser.NULL {
			txt := fmt.Sprintf("%v", val.Value)
			if len(txt) > 0 {
				fmt.Println(txt)
			}
		}
		i++
	}
	fmt.Println("Executed in", time.Since(start))
}
