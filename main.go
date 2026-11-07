package main

import (
	"os"
	"time"

	"github.com/brzx/rpwuko/repl"
)

func main() {
	time.Sleep(10 * time.Second)
	println("Hello user! This is the rpwuko programming language!")
	println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
