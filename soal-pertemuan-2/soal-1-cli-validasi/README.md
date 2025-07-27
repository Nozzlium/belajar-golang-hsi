# Age Verificator CLI

Age verificator paling oke di sokin

## Installation

simply

```bash
go run main.go 
```
di folder soal-2-http-validasi
## Usage
contoh:
```bash
Nama: Cesa
Umur: 18

# returns
Selamat datang, Cesa!
```
contoh input di bawah umur
```bash
Nama: Cesa
Umur: 12

# returns
Error: umur tidak valid (minimal 18 tahun)
```
contoh input nama kosong
```bash
Nama: 

# returns
Error: input tidak valid
```
contoh input umur kosong/bukan angka
```bash
Nama:       12
Umur: Aji

# returns
Error: input tidak valid
```