package repository

import (
    "database/sql"
    "HearingAidClinicClientManagement/pkg/model"
)

type PatientRepository struct {
    db *sql.DB
}

func NewPatientRepository(db *sql.DB) *PatientRepository {
    return &PatientRepository{db: db}
}

func (r *PatientRepository) Save(patient *model.Patient) error {
    _, err := r.db.Exec("INSERT INTO patients (id, name, age) VALUES ($1, $2, $3)", patient.ID, patient.Name, patient.Age)
    return err
}
