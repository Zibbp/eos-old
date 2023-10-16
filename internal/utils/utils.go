package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

func GetFileExtension(filePath string) string {
	fileSlice := strings.LastIndex(filePath, ".")
	fileExt := filePath[fileSlice+1:]
	return fileExt
}

func CreateDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	return nil
}

func RemoveDirectory(dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		return fmt.Errorf("failed to remove directory: %w", err)
	}
	return nil
}

func MoveFile(srcPath, dstPath string) error {
	// Open the source file for reading.
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file.
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy the contents of the source file to the destination file.
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// Close the destination file.
	if err := dstFile.Close(); err != nil {
		return err
	}

	// Remove the source file.
	if err := os.Remove(srcPath); err != nil {
		return err
	}

	return nil
}

func DownloadFile(url string, filePath string) error {
	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url) // nolint:gosec
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
