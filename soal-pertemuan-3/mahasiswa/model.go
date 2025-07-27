package mahasiswa

import "fmt"

type Deskripsi interface {
	Info() string
	RataRata() float64
	GetUmur() int
}

type Mahasiswa struct {
	Nama     string
	Nilai    []int
	umur     int
	nilaiAvg float64
}

func (mahasiswa *Mahasiswa) Info() string {
	return fmt.Sprintf("Nama: %s, Umur: %d", mahasiswa.Nama, mahasiswa.GetUmur())
}

func (mahasiswa *Mahasiswa) RataRata() float64 {
	return mahasiswa.nilaiAvg
}

func (mahasiswa *Mahasiswa) GetUmur() int {
	return mahasiswa.umur
}
