package utils

import (
	"github.com/google/uuid"
	"strings"
)

// GetOrCreateFileName get or generate file name from the context type
func GetOrCreateFileName(fileName string, contextType string) string {
	if fileName != "" {
		return fileName
	}

	typeArr := strings.Split(contextType, "/")
	if len(typeArr) < 2 || strings.Index("/image/video/audio", typeArr[0]) == -1 {
		return fileName
	}

	return uuid.New().String() + "." + typeArr[1]
}
