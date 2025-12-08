.intel_syntax noprefix

.data
nickname:
    .ascii "hafidz permana\n"

.text
    .global _start

_start:
    mov rax, 1
    mov rdi, 1
    lea rsi, [nickname]
    mov rdx, 15
    syscall

    mov rax, 60
    xor rdi, rdi
    syscall
