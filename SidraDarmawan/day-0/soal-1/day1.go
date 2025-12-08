package main

import (
	"syscall"
)

func main() {
	name := "Sidra Darmawan\n"
	data := []byte(name)
	syscall.Write(1, data)
}
