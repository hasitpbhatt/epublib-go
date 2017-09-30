package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/hasitpbhatt/epublib-go/constants"
	"github.com/hasitpbhatt/epublib-go/errors"
)

// GetExtension returns the extension of the file
func GetExtension(filename string) (string, error) {
	words := strings.Split(filename, ".")
	lastIndex := len(words) - 1
	if lastIndex == 0 {
		return "", errors.ErrNoExtFound
	}
	return words[lastIndex], nil
}

// IsValidFilename returns whether the filename is valid or not
func IsValidFilename(filename string) error {
	ext, err := GetExtension(filename)
	if err != nil {
		return fmt.Errorf(":Can't create book object: %v", errors.ErrNoExtFound)
	}

	if ext != constants.Extension {
		return errors.ErrInvalidExtension
	}

	_, err = os.Stat(filename)
	if err != nil {
		return errors.ErrFileDoesNotExist
	}
	return nil
}
