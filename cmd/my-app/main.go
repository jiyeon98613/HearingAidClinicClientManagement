package main

import (
	"fmt"
	"log"
	"net/http"

	"workspace/src/github.com/jiyeon98613/HearingAidClinicClientManagement/pkg/bootstrap"
	"workspace/src/github.com/jiyeon98613/HearingAidClinicClientManagement/pkg/model"
	"workspace/src/github.com/jiyeon98613/HearingAidClinicClientManagement/pkg/webhook"
)

func main() {
	// Initialize services
	patientService, err := bootstrap.InitializeServices()
	if err != nil {
		log.Fatalf("Failed to initialize services: %v", err)
	}

	// Add a new patient (sample operation)
	newPatient := &model.Patient{CID: "c_1", Name: "John Doe", Age: 30}
	err = patientService.AddPatient(newPatient)
	if err != nil {
		log.Fatalf("Error adding a new patient: %v", err)
	}
	fmt.Println("New patient added successfully!")

	// Set up HTTP server to listen for GitHub webhook events
	http.HandleFunc("/", webhook.HandleGitHubPushEvent)
	log.Println("Listening for webhooks on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
