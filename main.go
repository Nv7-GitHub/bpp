package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/Nv7-Github/Bpp/parser"
)

var filename string

func init() {
	flag.StringVar(&filename, "file", "", "File to run")
}

func handle(err error) {
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(0)
	}
}

func main() {
	flag.Parse()

	script, err := ioutil.ReadFile(filename)
	handle(err)
	src := strings.TrimSpace(string(script))

	start := time.Now()
	prog, err := parser.Parse(src)
	handle(err)
	fmt.Println("Parsed in", time.Since(start))

	start = time.Now()
	out, err := prog.Run()
	handle(err)
	fmt.Println("Executed in", time.Since(start))

	fmt.Println(strings.TrimSpace(out))
}
