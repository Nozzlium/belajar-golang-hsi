package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer
var reader *kafka.Reader

var registrationSuccessTopic = "student.registration_validated"
var topicStudentAcademicInitialized = "student.academic_initialized"

type Student struct {
	StudentID string `json:"studentId"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Domicile  string `json:"domicile"`
}

func main() {
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"127.0.0.1:9094"},
	})
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"127.0.0.1:9094"},
		GroupID: "academic-consumer-group",
		GroupTopics: []string{
			registrationSuccessTopic,
		},
	})
	defer func() {
		writer.Close()
		reader.Close()
	}()

	for {
		msg, _ := reader.ReadMessage(context.Background())
		var student *Student
		msgBytes := msg.Value
		_ = json.Unmarshal(msgBytes, &student)
		log.Printf("[AcademicService] Received event: %s", msg.Topic)
		if student.Status == "valid" {
			log.Printf("[AcademicService] Student academic successfully validated for student_id: %s", student.StudentID)
			_ = writer.WriteMessages(context.TODO(),
				kafka.Message{
					Key:   []byte(student.StudentID),
					Value: msgBytes,
					Topic: topicStudentAcademicInitialized,
				},
			)
			log.Printf("[AcademicService] Sent event: %s", topicStudentAcademicInitialized)
		}
	}
}
