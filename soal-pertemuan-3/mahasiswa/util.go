package mahasiswa

var maxNilai int = 100

var Versi string = "v1.0.0"

func hitungRataRata(nilai ...int) float64 {
	var sum float64 = 0
	var size float64 = float64(len(nilai))
	for _, val := range nilai {
		sum += float64(val)
	}
	return sum / size
}

func BuatMahasiswa(nama string, umur int, nilai ...int) Mahasiswa {
	return Mahasiswa{
		Nama:     nama,
		umur:     umur,
		Nilai:    nilai,
		nilaiAvg: hitungRataRata(nilai...),
	}
}

func GetMaxNilai() int {
	return maxNilai
}
