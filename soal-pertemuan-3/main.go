package main

import (
	"fmt"
	"tugas-pertemuan-3/mahasiswa"
)

func main() {
	ali := mahasiswa.BuatMahasiswa(
		"Ali",
		20,
		70, 100, 85,
	)
	printMahaiswa(&ali)

	siti := mahasiswa.BuatMahasiswa(
		"Siti",
		22,
		80, 85, 80,
	)
	printMahaiswa(&siti)

	fmt.Printf("Versi package: %s\n", mahasiswa.Versi)
	fmt.Printf("Nilai maksimum: %d\n", mahasiswa.GetMaxNilai())
	fmt.Printf("Total umur Mahasiswa: %d\n", getTotalAge(&ali, &siti))
}

func printMahaiswa(mhs *mahasiswa.Mahasiswa) {
	fmt.Printf("%s\nRata-rata nilai: %.2f\n---\n", mhs.Info(), mhs.RataRata())
}

func getTotalAge(mhs ...*mahasiswa.Mahasiswa) int {
	total := 0
	for _, m := range mhs {
		total += m.GetUmur()
	}
	return total
}
