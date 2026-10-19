package storage

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"relat/internal/platform/httperr"
)

type Storage struct {
	DestFolder string
}

func NewStorage(destFolder string) (*Storage, error) {
	err := folderExists(destFolder)
	if err != nil {
		return nil, err
	}

	return &Storage{DestFolder: destFolder}, nil
}

func folderExists(destFolder string) error {
	if _, err := os.Stat(destFolder); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(destFolder, os.ModePerm); err != nil {
			return httperr.NewHttpError("unable to create dest folder", http.StatusInternalServerError)
		}
	}

	return nil
}

func (s *Storage) formatFileName(fileName string) string {
	fileName = strings.ToLower(fileName)
	fileName = strings.ReplaceAll(fileName, " ", "-")

	return fmt.Sprintf("%s/%s", s.DestFolder, fileName)
}

func (s *Storage) SaveFile(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", httperr.NewHttpError("unable to open file content", http.StatusInternalServerError)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return "", httperr.NewHttpError("unable to read file content", http.StatusInternalServerError)
	}

	fileUrl := s.formatFileName(file.Filename)
	if err := os.WriteFile(fileUrl, bytes, os.ModePerm); err != nil {
		fmt.Println(err)
		return "", httperr.NewHttpError("unable to write file content", http.StatusInternalServerError)
	}

	return fileUrl, nil
}
