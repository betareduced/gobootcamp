;# This code only accepts under 4096 bytes
.intel_syntax noprefix

;# alokasi memori untuk variable buffer
.bss
.lcomm buffer, 4096

.text
.global _start
_start:
  ;# syscall_read
  mov rax, 0
  mov rdi, 0
  mov rsi, OFFSET buffer
  mov rdx, 4096
  syscall

  mov r10, rax

  ;# syscall_write
  mov rax, 1
  mov rdi, 1
  ;# since the rsi save the address and not changes
  mov rdx, 4096
  syscall

  ;# syscall_exit
  mov rax, 60
  mov rdi, 0
  syscall