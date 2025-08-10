package main

import (
	"fmt"
	"sync"

	"github.com/nozzlium/tugas-pertemuan-4/config"
	"github.com/nozzlium/tugas-pertemuan-4/models"
	"github.com/nozzlium/tugas-pertemuan-4/worker"
	"gorm.io/gorm"
)

func main() {
	db := config.GetDB()
	defer func() {
		if r := recover(); r != nil {

		}
		fmt.Println("truncating")
		truncate(db)
	}()

	db.AutoMigrate(
		&models.Mahasiswa{},
		&models.Tugas{},
		&models.Hasil{},
	)

	mahasiswas := []*models.Mahasiswa{
		{Nama: "Andi Pratama"},
		{Nama: "Budi Santoso"},
		{Nama: "Citra Lestari"},
		{Nama: "Dian Kusuma"},
		{Nama: "Eka Sari"},
	}

	jobMhs := make(chan *models.Mahasiswa)
	resultTugas := make(chan *models.Tugas)

	jobHasil := make(chan *models.Tugas)
	resultHasil := make(chan *models.Hasil)

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker.AssignmentWorker(db, jobMhs, resultTugas)
		go worker.GradingWorker(db, &wg, jobHasil, resultHasil)
	}

	_ = db.Create(mahasiswas)

	for _, mhs := range mahasiswas {
		jobMhs <- mhs
	}

	wg.Wait()

}

func truncate(db *gorm.DB) {
	db.Migrator().DropTable(&models.Hasil{}, &models.Tugas{}, &models.Mahasiswa{})
}
