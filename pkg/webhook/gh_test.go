package webhook

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGitHubPushEvent(t *testing.T) {
	requestBody := `{
		"ref": "refs/heads/main",
		"repository": {
			"name": "your-repo-name"
		},
		"pusher": {
			"name": "your-username"
		}
	}`
	req, err := http.NewRequest("POST", "/", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleGitHubPushEvent)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Event added to Google Calendar\n"
	if recorder.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", recorder.Body.String(), expected)
	}
}
