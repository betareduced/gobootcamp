.intel_syntax noprefix

.text
.global _start
_start:
  mov rax, 1
  mov rdi, 1
  lea rsi, nama
  mov rdx, 28
  syscall
  
  mov rax, 60
  xor rdi, rdi
  syscall

.data
nama:
  .ascii "Naufal Iffa Maulana Ramadhan"

