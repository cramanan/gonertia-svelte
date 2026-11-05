package main

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	getenv := func(key string) string {
		if key == "LISTEN_ADDR" {
			return ":18000"
		}
		return ""
	}
	errCh := make(chan error, 1)
	go func() { errCh <- run(ctx, getenv, io.Discard) }()

	var res *http.Response
	var err error
	for range 50 {
		if res, err = http.Get("http://localhost:18000/api/status"); err == nil {
			break
		}
		time.Sleep(20 * time.Millisecond)
	}
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK || !strings.Contains(string(body), `"status":"ok"`) {
		t.Fatalf("got %d %s", res.StatusCode, body)
	}

	// Unknown URLs render the Error page with the 404 kept, as HTML and as Inertia JSON.
	for _, inertia := range []string{"", "true"} {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:18000/nope", nil)
		req.Header.Set("X-Inertia", inertia)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		body, _ := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode != http.StatusNotFound || !strings.Contains(string(body), "Error") {
			t.Errorf("X-Inertia=%q: got %d %s", inertia, res.StatusCode, body)
		}
		if inertia != "" && res.Header.Get("X-Inertia") != "true" {
			t.Errorf("Inertia 404 lost its headers: %v", res.Header)
		}
	}
}
