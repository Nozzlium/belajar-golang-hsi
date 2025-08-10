package worker

import (
	"fmt"
	"sync"

	"github.com/nozzlium/tugas-pertemuan-4/models"
	"gorm.io/gorm"
)

var hasil = []*models.Hasil{
	{Nilai: 70},
	{Nilai: 80},
	{Nilai: 75},
	{Nilai: 79},
	{Nilai: 90},
}

func GradingWorker(
	db *gorm.DB,
	wg *sync.WaitGroup,
	assignmentChan chan<- *models.Tugas,
	resultChan chan<- *models.Hasil,
) {
	defer wg.Done()

	assignment := <-assignmentChan
	fmt.Println(assignment, "grading")
	grade := hasil[assignment.ID%5]
	grade.TugasID = assignment.ID
	result := db.Create(grade)
	err := result.Error
	if err == nil {
		fmt.Println("graded")
		resultChan <- grade
	}
}
