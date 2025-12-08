# Sidra Darmawan

ABI itu aturan main nya antara program dan sistem operasi. 

Kalau API pakai JSON, di ABI pakai “kontrak” antara program dan OS, yaitu syscall dipanggil dan data dikirim lewat register.

- Masukkan nomor syscall ke register khusus
- Masukkan argumen ke register lain rdi, rsi, rdx
- Panggil instruksi `syscall`

Register berbeda:  
- amd64 → rax untuk nomor syscall, rdi/rsi/rdx untuk argumen  
- arm64 → x8 untuk nomor syscall, x0/x1/x2 untuk argumen  

Program Go manggil `fmt.Println` → Go runtime atur string → runtime panggil syscall write → OS kernel tulis ke stdout → muncul di layar.