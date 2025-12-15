package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: goseq <start> <end>")
		os.Exit(1)
	}

	start, errStart := strconv.Atoi(os.Args[1])
	end, errEnd := strconv.Atoi(os.Args[2])

	if errStart != nil || errEnd != nil {
		fmt.Fprintln(os.Stderr, "goseq: invalid number argument")
		os.Exit(1)
	}

	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
}
