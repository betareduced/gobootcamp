package main

import (
	"syscall"
	"unsafe"
)

func main() {
	name := "Yonie Abdul Salam\n"
	data := []byte(name)
	// file descriptor to stdout is 1
	fd := 1
	// get address of first element
	buf := unsafe.Pointer(&data[0])
	//length of data
	lengthName := len(data)
	// call write using syscall
	syscall.Syscall(syscall.SYS_WRITE, uintptr(fd), uintptr(buf), uintptr(lengthName))
}
