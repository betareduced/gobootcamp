.intel_syntax noprefix

.text
.global _start
_start:
  mov rax, 1          
  mov rdi, 1          
  lea rsi, name
  mov rdx, 14         
  syscall             

  mov rax, 60         
  xor rdi, rdi        
  syscall

.data
name:
  .ascii "Sidra Darmawan\n"
