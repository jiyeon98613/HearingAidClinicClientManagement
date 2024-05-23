package repository

import (
    "database/sql"
    "HearingAidClinicClientManagement/pkg/model"
)

// PatientRepository describes the interface for patient storage operations
type PatientRepository interface {
    Save(patient *model.Patient) error
}

// MySQLPatientRepository is an implementation of PatientRepository using a MySQL database
type MySQLPatientRepository struct {
    db *sql.DB
}

// NewMySQLPatientRepository creates a new instance of MySQLPatientRepository
func NewMySQLPatientRepository(db *sql.DB) *MySQLPatientRepository {
    return &MySQLPatientRepository{db: db}
}

// Save inserts a new patient record into the database
func (r *MySQLPatientRepository) Save(patient *model.Patient) error {
    _, err := r.db.Exec("INSERT INTO patients (id, name, age) VALUES (?, ?, ?)", patient.ID, patient.Name, patient.Age)
    return err
}
