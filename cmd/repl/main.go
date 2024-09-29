package main

import (
	"fmt"
	"os"

	"github.com/olacin/monkey-interpreter/repl"
)

const usage = `Usage:
# Starting a classic REPL
./repl
# Interpreting a script
./repl path/to/script.monkey`

func main() {
	fmt.Println("Monkey REPL 0.0.1")

	// Classic REPL
	if len(os.Args) == 1 {
		repl.Start(os.Stdin, os.Stdout)
	} else if len(os.Args) == 2 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("unable to open file: %s", err)
		}
		repl.Start(file, os.Stdout)
	} else {
		fmt.Println(usage)
	}
}
