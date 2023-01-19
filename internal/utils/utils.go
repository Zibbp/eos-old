package utils

import (
	"fmt"
	"os"
)

func GetFoldersInDir(dir string) ([]string, error) {
	folders, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %w", err)
	}
	var folderNames []string
	for _, folder := range folders {
		if folder.IsDir() {
			folderNames = append(folderNames, folder.Name())
		}
	}
	return folderNames, nil
}

func GetFilesInDir(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %w", err)
	}
	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames, nil
}

func StringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
