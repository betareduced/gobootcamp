.intel_syntax noprefix

.text
.global _start
_start:
  ;# syscall_write
  mov rax, 1
  mov rdi, 1
  lea rsi, nama
  mov rdx, 18
  syscall

  ;# syscall_exit
  mov rax, 60
  mov rdi, 0
  syscall
  
.data
nama:
  .ascii "Yonie Abdul Salam\n"