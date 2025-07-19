package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var ErrMethodNotAllowd = errors.New("method not allowed")
var ErrInvalidAge = errors.New("umur tidak valid")
var ErrInvalidInput = errors.New("email kosong atau umur kurang dari 18")

type AgeService struct {
}

func NewAgeService() *AgeService {
	return &AgeService{}
}

func (service *AgeService) Validate(email string, age int) (string, error) {
	if email != "" && age >= 18 {
		return "ok", nil
	}
	log.WithFields(log.Fields{
		"location": "AgeService.Validate",
	}).Error(ErrInvalidInput.Error())
	return "", fmt.Errorf("%w", ErrInvalidInput)
}

type AgeController struct {
	AgeService *AgeService
}

func NewAgeController(ageService *AgeService) *AgeController {
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
		http.Error(w, ErrMethodNotAllowd.Error(), http.StatusMethodNotAllowed)
	}

	email := r.URL.Query().Get("email")
	ageString := r.URL.Query().Get("age")
	age, err := strconv.Atoi(ageString)
	if err != nil {
		http.Error(w, ErrInvalidAge.Error(), http.StatusBadRequest)
	}

	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	status, err := controller.AgeService.Validate(email, age)
	if err != nil && errors.Is(err, ErrInvalidInput) {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(Response{
			Error: "email kosong atau umur kurang dari 18",
		})
		return
	}
	encoder.Encode(Response{
		Status: status,
	})
}

func main() {
	ageService := NewAgeService()
	ageController := NewAgeController(ageService)

	http.HandleFunc("/validate", ageController.Validate)
	http.ListenAndServe(":8080", nil)
}
