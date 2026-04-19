package helper

import (
	"fmt"
	"image_storage/config"
	"mime/multipart"
	"time"
)

func UploadImage(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	timestamp := time.Now().Format("20060102150405")
	ext := ""
	for i := len(fileHeader.Filename) - 1; i >= 0; i-- {
		if fileHeader.Filename[i] == '.' {
			ext = fileHeader.Filename[i:]
			break
		}
	}
	filename := fmt.Sprintf("%s%s", timestamp, ext)

	_, err = config.Supabase.Storage.UploadFile(config.SupabaseBucket, filename, file)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func DeleteImage(path string) error {
	if path == "" {
		return nil
	}

	_, err := config.Supabase.Storage.RemoveFile(
		config.SupabaseBucket,
		[]string{path},
	)

	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	return nil
}
