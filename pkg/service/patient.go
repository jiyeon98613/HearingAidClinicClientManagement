package service

import (
	"github.com/jiyeon98613/HearingAidClinicClientManagement/pkg/model"
	"github.com/jiyeon98613/HearingAidClinicClientManagement/pkg/repository"
)

// PatientService provides high-level logic for patient data
type PatientService struct {
	repo repository.PatientRepository
}

// NewPatientService creates a new PatientService
func NewPatientService(repo repository.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

// AddPatient adds a new patient
func (s *PatientService) AddPatient(patient *model.Patient) error {
	return s.repo.Save(patient)
}
