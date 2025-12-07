.intel_syntax noprefix
.global _start

.section .data
msg:
    .ascii "Muhamad Rizky Athoriq\n"
msg_len = . - msg

.section .text
_start:
    mov rax, 1
    mov rdi, 1
    lea rsi, [rip + msg]
    mov rdx, msg_len
    syscall

    mov rax, 60
    xor rdi, rdi
    syscall

