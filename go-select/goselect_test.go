package goselect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func createMockServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {

	t.Run("Should return the fastest URLs when both servers respond within time", func(t *testing.T) {
		slowServer := createMockServerWithDelay(20 * time.Millisecond)
		fastServer := createMockServerWithDelay(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Should return timeout error if a server takes more than 10 seconds", func(t *testing.T) {
		server := createMockServerWithDelay(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}

	})
}
