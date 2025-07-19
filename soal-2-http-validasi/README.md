# Age Verificator HTTP

Age verificator nomor satu di sokin

## Installation

simply

```bash
go run main.go 
```
di folder soal-2-http-validasi
## Usage
contoh:
```bash
curl "http://localhost:8080/validate?email=test@example.com&age=20"
# returns
{"status":"ok"}
```
```bash
curl "http://localhost:8080/validate?email=test@example.com&age=20"
# returns
{"error":"email kosong atau umur kurang dari 18"}
```