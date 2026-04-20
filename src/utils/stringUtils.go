package utils

import (
	"fmt"
	"path/filepath"
	"strings"

	customerrors "github.com/yepakh/notepad/src/customErrors"
)

func IsValidPathOrEmpty(path string) error {
	if strings.ContainsAny(path, "\x00") {
		return fmt.Errorf("path contains null byte")
	}

	clean := filepath.Clean(path)

	if path == "" || clean == "." {
		return fmt.Errorf("'%v' %v", path, customerrors.ErrEmptyPath)
	}

	return nil
}
