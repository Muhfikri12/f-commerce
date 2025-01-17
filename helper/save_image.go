package helper

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context, formFieldName string) (string, error) {

	file, err := c.FormFile(formFieldName)
	if err != nil {
		return "", errors.New("failed to retrieve file: " + err.Error())
	}

	if !isImage(file.Filename) {
		return "", errors.New("invalid file type, only images are allowed")
	}

	uploadPath := "./uploads/"

	if err := ensureDir(uploadPath); err != nil {
		return "", errors.New("failed to create upload directory: " + err.Error())
	}

	filePath := filepath.Join(uploadPath, file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", errors.New("failed to save file: " + err.Error())
	}

	return filePath, nil
}

func isImage(filename string) bool {
	validExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	ext := strings.ToLower(filepath.Ext(filename))
	for _, validExt := range validExtensions {
		if ext == validExt {
			return true
		}
	}
	return false
}

func ensureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}
