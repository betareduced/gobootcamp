### Questions

1. Apa itu ABI?
2. Kalau di API kita menggunakan request body, seperti json part dari contract yang kita gunakan untuk berinteraksi dengan backend. Kalau di ABI, apa contract yang kita gunakan?
3. Bagaimana cara memanggil syscall di assembly?
4. Apakah yang membedakan syscall di arm64 dan syscall di amd64?
5. Coba jelaskan dengan bahasa kalian sendiri, apa yang terjadi dibalik program ini dari high level-hingga ke low level:
```go
fmt.Println("Hello, World!")
```

### Answers
1. ABI adalah aturan bahasa mesin yang memastikan program assembly dapat berinteraksi dengan hardware sesuai dengan contract yang telah ditentukan
2. contract yang digunakan harus sesuai kernel mode yang ditentukan dan register yang akan dipanggil harus sesuai dengan systemcall yang akan digunakan
3. Cara memanggil `syscall` di assembly adalah :
    - pilih kernel mode yang akan digunakan, (misalnya `x86-64`, `arm64`, dst.)
    - pilih `syscall` instruction sesuai dengan kernel mode yang digunakan
    - ikuti aturan contract seperti memasukkan `argumen/parameter` sesuai `syscall` yang akan dipanggil
4.  yang membedakan `syscall` di arm64 dan `syscall` di amd64 adalah:
    - register yang digunakan untuk `return value` maupun `argumen/parameter` berbeda
5. `fmt.Println("Hello, World!")` dijelaskan dari high-level ke low-level :
    - Dengan asumsi, program tersebut telah di compile dalam bentuk executable, Program golang akan me
    - `fmt` adalah standard library dari Golang untuk formating.
    - di dalam library `fmt`, ada method `Println` untuk mencetak data ke standard output.
    - Pada method `Println`, ia akan memanggil fungsi `write` menggunakan `syscall` menyesuaikan dengan data yang di dalam argument tersebut.
    - Di dalam Assembly, instruksi `syscall write` akan mengeksekusi sesuai dengan `register write` yaitu 1 dan file descriptor `stdout` yaitu 1.
    - Kernel akan menerima yang instruksi tersebut dan menampilkan pada terminal/console.
