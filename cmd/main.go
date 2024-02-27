package main

import (
	"fmt"
	"os"

	"github.com/olacin/monkey-interpreter/repl"
)

func main() {
	fmt.Println("Monkey REPL 0.0.1")
	repl.Start(os.Stdin, os.Stdout)
}
