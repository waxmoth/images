package auth

import (
	"image-functions/src/services/auth"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	jwtService := auth.JWTService{Key: "Test_auth_key", Expires: 72}
	data := "Test_auth_data"

	tokenString, err := jwtService.Encode(data)
	if err != nil {
		t.Fatalf("Failed to encode token, err: %v", err)
	}
	decodedData, err := jwtService.Decode(tokenString)
	if err != nil {
		t.Fatalf("Failed to decode token, err: %v", err)
	}
	if decodedData != data {
		t.Fatalf("Expected %s, but got %s", data, decodedData)
	}
}

func BenchmarkEncodeData(b *testing.B) {
	jwtService := auth.JWTService{Key: "Test_auth_key", Expires: 72}
	for i := 0; i < b.N; i++ {
		_, err := jwtService.Encode("Test_auth_data")
		if err != nil {
			return
		}
	}
}

func BenchmarkDecodeData(b *testing.B) {
	jwtService := auth.JWTService{Key: "Test_auth_key", Expires: 72}
	token, _ := jwtService.Encode("Test_auth_data")
	for i := 0; i < b.N; i++ {
		_, err := jwtService.Decode(token)
		if err != nil {
			return
		}
	}
}
