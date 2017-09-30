package book

import (
	"archive/zip"
	"io"

	"github.com/hasitpbhatt/epublib-go/errors"

	"github.com/hasitpbhatt/epublib-go/constants"
	"github.com/hasitpbhatt/epublib-go/utils"
)

// Book is common struct for all ebook formats
type Book struct {
	// Format indicates the ebook format
	Format string

	//
	pathFileMap map[string]*zip.File
}

// NewBookFromFile creates ebook object of the book
func NewBookFromFile(filename string) (*Book, error) {

	err := utils.IsValidFilename(filename)

	if err != nil {
		return nil, err
	}

	reader, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}

	pathFileMap := make(map[string]*zip.File)

	for _, file := range reader.File {
		pathFileMap[file.Name] = file
	}

	return &Book{Format: constants.Extension, pathFileMap: pathFileMap}, nil
}

// NewBook creates empty object
func NewBook() (*Book, error) {
	return &Book{}, nil
}

// GetPaths returns the list of filepaths in the epub
func (b *Book) GetPaths() []string {
	paths := []string{}
	for path := range b.pathFileMap {
		paths = append(paths, path)
	}
	return paths
}

// Open opens a specific file and returns ReadCloser object
func (b *Book) Open(filepath string) (io.ReadCloser, error) {
	if _, ok := b.pathFileMap[filepath]; !ok {
		return nil, errors.ErrFileDoesNotExist
	}
	file := b.pathFileMap[filepath]
	return file.Open()
}
