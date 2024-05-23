package bootstrap

import (
	"github.com/jiyeon98613/HearingAidClinicClientManagement/pkg/database"
	"github.com/jiyeon98613/HearingAidClinicClientManagement/pkg/repository"
	"github.com/jiyeon98613/HearingAidClinicClientManagement/pkg/service"
)

// InitializeServices sets up the database connection and repositories
func InitializeServices() (*service.PatientService, error) {
	// Set up the database connection
	dsn := "root:root@tcp(localhost:3306)/hearingaid_center"
	db := database.InitDB(dsn)

	// Initialize the repositories
	patientRepo := repository.NewMySQLPatientRepository(db)

	// Initialize the services
	patientService := service.NewPatientService(patientRepo)

	return patientService, nil
}
