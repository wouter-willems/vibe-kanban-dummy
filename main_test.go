package main

import (
	"io"
	"net/http"
	"testing"
	"time"
)

func get(t *testing.T, url string) (int, string) {
	t.Helper()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		t.Fatalf("GET %s: %v", url, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	return resp.StatusCode, string(body)
}

func TestLiveServer(t *testing.T) {
	cases := []struct {
		url  string
		want string
	}{
		{"http://localhost:8080/hello", `{"message":"hello, world"}` + "\n"},
		{"http://localhost:8080/hello?name=Wouter", `{"message":"hello, Wouter"}` + "\n"},
	}
	for _, tc := range cases {
		status, body := get(t, tc.url)
		if status != 200 {
			t.Errorf("%s: status = %d, want 200", tc.url, status)
		}
		if body != tc.want {
			t.Errorf("%s: body = %q, want %q", tc.url, body, tc.want)
		}
	}
}
