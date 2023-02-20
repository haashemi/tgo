package tgo

import (
	"io"
)

type InputFile interface {
	NeedsUpload() bool
}

type InputFileUploadable struct {
	Name   string
	Reader io.Reader
}

type InputFileNotUploadable string

func (InputFileUploadable) NeedsUpload() bool { return true }

func (InputFileNotUploadable) NeedsUpload() bool { return false }

func FileFromID(fileID string) InputFile {
	return InputFileNotUploadable(fileID)
}

func FileFromURL(url string) InputFile {
	return InputFileNotUploadable(url)
}

func FileFromReader(name string, reader io.Reader) InputFile {
	return &InputFileUploadable{Name: name, Reader: reader}
}
