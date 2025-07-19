package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nozzlium/belajar-golang-hsi/soal-2-http-validasi/service"
)

type AgeController struct {
	AgeService *service.AgeService
}

func NewAgeController(ageService *service.AgeService) *AgeController {
	return &AgeController{
		AgeService: ageService,
	}
}

func (controller *AgeController) Validate(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string `json:"status,omitempty"`
		Error  string `json:"error,omitempty"`
	}

	if method := r.Method; method != "GET" {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}

	email := r.URL.Query().Get("email")
	ageString := r.URL.Query().Get("age")
	age, err := strconv.Atoi(ageString)
	if err != nil {
		http.Error(w, "invalid age format", http.StatusBadRequest)
	}

	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	if ok := controller.AgeService.Validate(email, age); ok {
		encoder.Encode(Response{
			Status: "ok",
		})
		return
	}
	encoder.Encode(Response{
		Error: "email kosong atau umur kurang dari 18",
	})
}
