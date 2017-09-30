package errors

import "errors"

var (
	// ErrNoExtFound indicates the extension of the filename could not be figured
	ErrNoExtFound = errors.New("No extension was found")

	// ErrInvalidExtension indicates the extension of the filename is not valid
	ErrInvalidExtension = errors.New("Invalid extension")

	// ErrFileDoesNotExist indicates the file does not exist
	ErrFileDoesNotExist = errors.New("The file does not exist")
)
