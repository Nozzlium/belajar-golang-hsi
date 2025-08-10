package worker

import (
	"fmt"

	"github.com/nozzlium/tugas-pertemuan-4/models"
	"gorm.io/gorm"
)

var tugas = []*models.Tugas{
	{Judul: "Tugas Pemrograman Goroutine", Deskripsi: "Tugas Pemrograman Goroutine"},
	{Judul: "Tugas Implementasi WaitGroup", Deskripsi: "Tugas Implementasi WaitGroup"},
	{Judul: "Tugas Implementasi Mutex", Deskripsi: "Tugas Implementasi Mutex"},
	{Judul: "Tugas Implementasi Channel", Deskripsi: "Tugas Implementasi Channel"},
	{Judul: "Tugas Remidial Implementasi WaitGroup", Deskripsi: "Tugas Remidial Pemrograman Goroutine"},
}

func AssignmentWorker(
	db *gorm.DB,
	studentChan <-chan *models.Mahasiswa,
	assigned <-chan *models.Tugas,
) {
	student := <-studentChan
	fmt.Println(student, "assiging")
	assignment := tugas[student.ID%5]
	assignment.MahasiswaID = student.ID
	result := db.Create(assignment)
	err := result.Error
	fmt.Println(err)
	if err == nil {
		fmt.Println("assigned")
		assigned <- assignment
	}
}
