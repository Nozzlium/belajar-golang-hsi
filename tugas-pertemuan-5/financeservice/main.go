package main

import (
	"context"
	"encoding/json"
	"log"
	"slices"

	"github.com/segmentio/kafka-go"
)

var topicStudentRegistered = "student.registered"
var registrationSuccessTopic = "student.registration_validated"
var registrationFailedTopic = "student.registration_failed"

type Student struct {
	StudentID string `json:"studentId"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Domicile  string `json:"domicile"`
}

var writer *kafka.Writer
var reader *kafka.Reader

var blacklist []string = []string{
	"Depok",
}

func main() {
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9094"},
	})
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9094"},
		GroupID: "finance-consumer-group",
		GroupTopics: []string{
			topicStudentRegistered,
		},
	})
	defer func() {
		writer.Close()
		reader.Close()
	}()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("[FinanceService] Error: %v\n", err)
		}
		log.Printf("[FinanceService] Received event: %s\n", msg.Topic)
		switch msg.Topic {
		case topicStudentRegistered:
			var student *Student
			json.Unmarshal(msg.Value, &student)

			if slices.Contains(blacklist, student.Domicile) {
				log.Printf("[Finance Service] Payment failed to validate for student_id: %s", student.StudentID)
				student.Status = "invalid"
				invalidMsg, _ := json.Marshal(&student)
				write(writer, []byte(student.StudentID), invalidMsg, registrationFailedTopic)
				log.Printf("[Finance Service] Sent event: %s", registrationFailedTopic)
			} else {
				log.Printf("[Finance Service] Payment validated for student_id: %s", student.StudentID)
				student.Status = "valid"
				validMsg, _ := json.Marshal(&student)
				write(writer, []byte(student.StudentID), validMsg, registrationSuccessTopic)
				log.Printf("[Finance Service] Sent event: %s", registrationSuccessTopic)
			}
		}
	}
}

func write(writer *kafka.Writer, key []byte, msg []byte, topic string) {
	_ = writer.WriteMessages(context.TODO(),
		kafka.Message{
			Key:   key,
			Value: msg,
			Topic: topic,
		},
	)
}
