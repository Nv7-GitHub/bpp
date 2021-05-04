package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Nv7-Github/Bpp/parser"
	"github.com/davecgh/go-spew/spew"
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

	out, err := os.Create("ast.txt")
	handle(err)
	spew.Fdump(out, prog)
}
