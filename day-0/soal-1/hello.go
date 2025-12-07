package main

import (
	"syscall"
	"unsafe"
)

func main() {
	nickname := "hafidz permana\n"
	data := []byte(nickname)
	fd := 1
	buf := unsafe.Pointer(&data[0])
	lengthName := len(data)

	// write
	// syscall.Syscall(trap, a1, a2, a3 uintptr)
	// trap = syscall number
	// a1 = rdi = fd
	// a2 = rsi = buffer pointer
	// a3 = rdx = buffer length
	syscall.Syscall(
		syscall.SYS_WRITE,   // rax = 1
		uintptr(fd),         // rdi = 1
		uintptr(buf),        // rsi = pointer ke buffer
		uintptr(lengthName), // rdx = 15
	)

	// exit
	// trap = syscall number
	// a1 = rdi = exit code
	syscall.Syscall(
		syscall.SYS_EXIT, // rax = 60
		0,                // rdi = 0
		0,                // unuse
		0,                //
	)
}
