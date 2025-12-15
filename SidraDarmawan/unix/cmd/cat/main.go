package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
		}
		return
	}

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "gocat: cannot open file:", filename, err)
			continue
		}

		if _, err = io.Copy(os.Stdout, file); err != nil {
			fmt.Fprintln(os.Stderr, "gocat: error reading file:", filename, err)
		}
		file.Close()
	}
}
