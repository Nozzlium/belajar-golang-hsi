package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer
var reader *kafka.Reader

var topicStudentRegistered = "student.registered"
var topicStudentAcademicInitialized = "student.academic_initialized"
var topicStudentRegistrationFailed = "student.registration_failed"

func consumeMessage(reader *kafka.Reader) {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("[StudentService] Error: %v\n", err)
		}
		log.Printf("[StudentService] Received event: %s\n", msg.Topic)
		var student *Student
		_ = json.Unmarshal(msg.Value, &student)
		switch msg.Topic {
		case topicStudentRegistrationFailed:
			log.Printf("[StudentService] Failed to register for student_id: %s", student.StudentID)
		case topicStudentAcademicInitialized:
			log.Printf("[StudentService] Successfully registered student for student_id: %s", student.StudentID)
		}
	}
}

type Student struct {
	StudentID string `json:"studentId"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Domicile  string `json:"domicile"`
}

type RegisterController struct{}

func (controller *RegisterController) Register(w http.ResponseWriter, r *http.Request) {
	var student Student
	bodyBytes, _ := io.ReadAll(r.Body)
	json.Unmarshal(bodyBytes, &student)

	_ = writer.WriteMessages(context.TODO(),
		kafka.Message{
			Key:   []byte(student.StudentID),
			Value: bodyBytes,
			Topic: topicStudentRegistered,
		},
	)
	log.Printf("[StudentService] Sent event: %s", topicStudentRegistered)
}

func main() {
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9094"},
	})
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9094"},
		GroupID: "student-consumer-group",
		GroupTopics: []string{
			topicStudentRegistrationFailed,
			topicStudentAcademicInitialized,
		},
	})
	defer func() {
		writer.Close()
		reader.Close()
	}()

	go consumeMessage(reader)

	studentController := &RegisterController{}
	http.HandleFunc("/register", studentController.Register)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}

}
