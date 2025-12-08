package main

import (
	"os"
	"syscall"
	"unsafe"
)

func main() {
    full_name := []byte("Muhamad Rizky Athoriq\n")

    ptr := uintptr(unsafe.Pointer(&full_name[0]))

    syscall.Syscall(syscall.SYS_WRITE,
        uintptr(os.Stdout.Fd()),
        ptr,
        uintptr(len(full_name)),
    )
}
