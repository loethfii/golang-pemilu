package components

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func UpFileToLocal(file *multipart.FileHeader) error {
	// Open the uploaded file.
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	
	// Create the upload folder path if it doesn't exist.
	uploadFolder := filepath.Join(dir, "internal/upload")
	if _, err := os.Stat(uploadFolder); os.IsNotExist(err) {
		if err := os.Mkdir(uploadFolder, 0755); err != nil {
			return err
		}
	}
	
	fileName := fmt.Sprintf("%s-%s", time.Now().Format("20060102150405"), file.Filename)
	// Create a destination file for the uploaded content.
	pathImage := filepath.Join(uploadFolder, fileName)
	dst, err := os.Create(pathImage)
	if err != nil {
		return err
	}
	defer dst.Close()
	
	// Copy the uploaded content to the destination file.
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	
	return nil
}
