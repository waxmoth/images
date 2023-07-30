package tests

import (
	"fmt"
	"image-functions/src/consts"
	"image-functions/src/routers"
	"image-functions/src/services/auth"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestImageAPIReturnError(t *testing.T) {
	ts := httptest.NewServer(routers.Routers())
	defer ts.Close()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/image", ts.URL), nil)

	if err != nil {
		t.Fatalf("Expected no error on set request, got %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Expected no error on get response, got %v", err)
	}

	if resp.StatusCode != 401 {
		t.Fatalf("Expected status code 401, got %v", resp.StatusCode)
	}

	setAuthHeader(req, "bad key")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Expected no error on get response, got %v", err)
	}
	if resp.StatusCode != 401 {
		t.Fatalf("Expected get status code 401 when set wrong screct key, got %v", resp.StatusCode)
	}

	if os.Getenv("AUTH_KEY") == "" {
		os.Setenv("AUTH_KEY", "default_auth_key")
	}

	setAuthHeader(req, os.Getenv("AUTH_KEY"))
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Expected no error on get response, got %v", err)
	}
	if resp.StatusCode != 400 {
		t.Fatalf("Expected get status code 400, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}

func BenchmarkImageAPI(b *testing.B) {
	ts := httptest.NewServer(routers.Routers())
	defer ts.Close()
	for i := 0; i < b.N; i++ {
		_, err := http.Get(fmt.Sprintf("%s/image", ts.URL))
		if err != nil {
			return
		}
	}
}

func setAuthHeader(req *http.Request, key string) {
	req.Header.Set(consts.AuthUser, "test")
	if key != "" {
		authService := auth.JWTService{Key: key, Expires: 72}
		token, err := authService.Encode("a mocked data")
		if err != nil {
			panic(err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
	} else {
		req.Header.Set("Authorization", "Bearer bad_token")
	}
}
