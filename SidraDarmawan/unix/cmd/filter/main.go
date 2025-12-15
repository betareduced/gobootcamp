package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: gofilter <odd|even>")
		os.Exit(1)
	}

	mode := strings.ToLower(os.Args[1])
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		m, err := strconv.Atoi(text)
		if err != nil {
			continue
		}

		if mode == "odd" && m%2 == 1 {
			fmt.Println(m)
		}

		if mode == "even" && m%2 == 0 {
			fmt.Println(m)
		}
	}
}
