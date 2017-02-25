package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aprice/gordinals/ordinal"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "You must pass exactly one integer argument.")
		os.Exit(2)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s is not an integer.", os.Args[1])
		os.Exit(2)
	}
	fmt.Print(ordinal.For(n))
	os.Exit(0)
}
