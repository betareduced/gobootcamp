package main

import(
	"syscall"
    "unsafe"
)

func main(){
	nama := "Naufal Iffa Maulana Ramadhan\n"
    
    fd :=uintptr(1) 
    buf := uintptr(unsafe.Pointer(&[]byte(nama)[0]))
    lenght := uintptr(len(nama))
    
    syscall.Syscall(
    	syscall.SYS_WRITE,
        fd,
        buf,
        lenght,
     )
}