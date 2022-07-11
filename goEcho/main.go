package main

import (
	"fmt"
	"os"
)

func main() {
	var args []string = os.Args
	for _, arg := range args[1:] {
		fmt.Fprintf(os.Stdout, arg + " ")
	}
	fmt.Fprintf(os.Stdout, "\n")
}