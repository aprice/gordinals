package main

import (
	"fmt"
	"os"
	"strconv"
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
	switch n % 10 {
	case 1:
		fmt.Print("st")
	case 2:
		fmt.Print("nd")
	case 3:
		fmt.Print("rd")
	default:
		fmt.Print("th")
	}
	os.Exit(0)
}
