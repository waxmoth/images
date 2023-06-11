package utils

import (
	"image-functions/src/utils"
	"testing"
)

func TestGetOrCreateFileName(t *testing.T) {
	fileName := "abc.jpg"
	newFileName := utils.GetOrCreateFileName(fileName, "")
	if newFileName != fileName {
		t.Fatalf("Expected get file name %s, but got %s", fileName, newFileName)
	}

	if utils.GetOrCreateFileName("", "image/jpeg") == "" {
		t.Fatalf("Expected can create image file name")
	}

	if utils.GetOrCreateFileName("", "text/plain") != "" {
		t.Fatalf("Expected not create file name for text/plain")
	}
}

func BenchmarkCreateUUIDFileName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.GetOrCreateFileName("", "image/jpeg")
	}
}

func BenchmarkGetFileName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.GetOrCreateFileName("abc.txt", "image/jpeg")
	}
}
