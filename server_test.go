package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServe(t *testing.T) {
	server := httptest.NewServer(newServeMux())
	defer server.Close()

	t.Log(server)

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
}
