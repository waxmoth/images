package tests

import (
	"fmt"
	"image-functions/src/routers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestImageAPIReturnError(t *testing.T) {
	ts := httptest.NewServer(routers.Routers())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/image", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 400 {
		t.Fatalf("Expected status code 400, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}
