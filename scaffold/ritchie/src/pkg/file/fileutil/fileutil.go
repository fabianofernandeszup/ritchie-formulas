package fileutil

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Exists check if file exists
func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateIfNotExists creates dir if not exists
func CreateIfNotExists(dir string) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

// ReadFile wrapper for ioutil.ReadFile
func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// WriteFile wrapper for ioutil.WriteFile
func WriteFile(path string, content []byte) error {
	if !Exists(path) {
		_ , err := os.Create(path)
		if err != nil {
			panic(err)
		}
	}
	return ioutil.WriteFile(path, content, 0644)
}