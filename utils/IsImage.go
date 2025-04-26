package utils

import (
	"mime/multipart"
	"net/http"
	"strings"
)



func IsImage(file multipart.File) bool {
	buf := make([]byte, 512) 
	_, err := file.Read(buf)

	if err != nil {
		return false
	}

	_, err = file.Seek(0,0)
	if err != nil {
		return false
	}

	mimeType := http.DetectContentType(buf)

	return strings.HasPrefix(mimeType, "image/")
}