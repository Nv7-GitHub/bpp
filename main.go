package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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

	start := time.Now()
	executable, err := parser.Parse(string(script))
	handle(err)
	fmt.Println("Parsed in", time.Since(start))

	for _, val := range executable.Program {
		ret, err := val.Exec(executable)
		handle(err)
		if ret.Type != parser.NULL {
			fmt.Println(ret.Data)
		}
	}
}
