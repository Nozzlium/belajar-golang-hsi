package main

import (
	"net/http"

	"github.com/nozzlium/belajar-golang-hsi/soal-2-http-validasi/controller"
	"github.com/nozzlium/belajar-golang-hsi/soal-2-http-validasi/service"
)

func main() {
	ageService := service.NewAgeService()
	ageController := controller.NewAgeController(ageService)

	http.HandleFunc("/validate", ageController.Validate)
	http.ListenAndServe(":8080", nil)
}
