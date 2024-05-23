package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// GitHubPushEvent defines the structure for GitHub push event JSON payload
type GitHubPushEvent struct {
	Ref        string `json:"ref"`
	Repository struct {
		Name string `json:"name"`
	} `json:"repository"`
	Pusher struct {
		Name string `json:"name"`
	} `json:"pusher"`
}

// AddEventToCalendar creates a new event in the user's Google Calendar
func AddEventToCalendar(summary, description string) error {
	ctx := context.Background()

	calendarService, err := calendar.NewService(ctx, option.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
	if err != nil {
		return fmt.Errorf("unable to create Calendar service: %v", err)
	}

	event := &calendar.Event{
		Summary:     summary,
		Description: description,
		Start: &calendar.EventDateTime{
			DateTime: "2024-05-20T09:00:00-07:00", // Example start time
			TimeZone: "America/Los_Angeles",
		},
		End: &calendar.EventDateTime{
			DateTime: "2024-05-20T10:00:00-07:00", // Example end time
			TimeZone: "America/Los_Angeles",
		},
	}

	calendarID := "primary" // 'primary' calendar is default
	event, err = calendarService.Events.Insert(calendarID, event).Do()
	if err != nil {
		return fmt.Errorf("unable to create event: %v", err)
	}
	return nil
}

// HandleGitHubPushEvent processes incoming GitHub push events
func HandleGitHubPushEvent(w http.ResponseWriter, r *http.Request) {
	var event GitHubPushEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	summary := fmt.Sprintf("New push to %s", event.Repository.Name)
	description := fmt.Sprintf("Pushed by %s to branch %s", event.Pusher.Name, event.Ref)

	if err := AddEventToCalendar(summary, description); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Event added to Google Calendar")
}
