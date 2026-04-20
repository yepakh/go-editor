package utils

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

var ErrEmptyPath = errors.New("path is empty")

func IsValidPathOrEmpty(path string) error {
	if strings.ContainsAny(path, "\x00") {
		return fmt.Errorf("path contains null byte")
	}

	clean := filepath.Clean(path)

	if path == "" || clean == "." {
		return ErrEmptyPath
	}

	return nil
}
