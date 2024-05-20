package main

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "os"

    "google.golang.org/api/calendar/v3"
    "google.golang.org/api/option"
)

// GitHub Push Event 구조체 정의
type GitHubPushEvent struct {
    Ref        string `json:"ref"`
    Repository struct {
        Name string `json:"name"`
    } `json:"repository"`
    Pusher struct {
        Name string `json:"name"`
    } `json:"pusher"`
    // 필요한 필드 추가
}

// Google Calendar에 이벤트 추가 함수
func addEventToCalendar(summary, description string) error {
    ctx := context.Background()
    calendarService, err := calendar.NewService(ctx, option.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
    if err != nil {
        return fmt.Errorf("unable to create Calendar service: %v", err)
    }

    event := &calendar.Event{
        Summary:     summary,
        Description: description,
        Start: &calendar.EventDateTime{
            DateTime: "2024-05-20T09:00:00-07:00", // 임의의 시간 설정
            TimeZone: "America/Los_Angeles",
        },
        End: &calendar.EventDateTime{
            DateTime: "2024-05-20T10:00:00-07:00", // 임의의 시간 설정
            TimeZone: "America/Los_Angeles",
        },
    }

    calendarID := "primary" // 기본 캘린더 사용
    event, err = calendarService.Events.Insert(calendarID, event).Do()
    if err != nil {
        return fmt.Errorf("unable to create event: %v", err)
    }
    return nil
}

// HTTP 핸들러 함수
func handlePushEvent(w http.ResponseWriter, r *http.Request) {
    var event GitHubPushEvent
    if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Google Calendar에 이벤트 추가
    summary := fmt.Sprintf("New push to %s", event.Repository.Name)
    description := fmt.Sprintf("Pushed by %s to branch %s", event.Pusher.Name, event.Ref)
    if err := addEventToCalendar(summary, description); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintln(w, "Event added to Google Calendar")
}

func main() {
    http.HandleFunc("/", handlePushEvent)
    http.ListenAndServe(":8080", nil)
}

