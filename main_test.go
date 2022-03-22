package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	got := Ping()
	want := "PONG"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestBing(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	bing(response, request)

	t.Run("testing apollos apis", func(t *testing.T) {
		got := response.Body.String()
		want := "bong"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

}
