package main

import (
    "database/sql"
    "fmt"
    "log"
    "github.com/jiyeon98613/HearingAidClinicClientManagement/pkg/model"
    "github.com/jiyeon98613/HearingAidClinicClientManagement/pkg/repository"
    //"HearingAidClinicClientManagement/pkg/service"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // MySQL 연결을 설정합니다. 여기서는 root 사용자와 해당 패스워드를 사용합니다.
    // 'root'와 'root_password'를 실제 root 사용자의 이름과 패스워드로 변경해야 합니다.
    // 'my_database'는 연결하려는 데이터베이스 이름입니다.
    dsn := "root:root@tcp(localhost:3306)/hearingaid_center"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Error connecting to the database: ", err)
    }
    defer db.Close()

    // 데이터베이스 연결 테스트
    err = db.Ping()
    if err != nil {
        log.Fatal("Error pinging the database: ", err)
    }

    // Repository, Service 초기화
    patientRepo := repository.NewMySQLPatientRepository(db)
    patientService := service.NewPatientService(patientRepo)

    // 새로운 환자 정보를 추가
    newPatient := &model.Patient{ID: 1, Name: "John Doe", Age: 30}
    err = patientService.AddPatient(newPatient)
    if err != nil {
        log.Fatal("Error adding a new patient: ", err)
    }

    fmt.Println("New patient added successfully!")
}

