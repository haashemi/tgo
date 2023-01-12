package tgo

import "io"

type InputFile interface{ NeedsUpload() bool }

type InputFileNoUpload string

func (InputFileNoUpload) NeedsUpload() bool { return false }

func FileFromID(fileID string) InputFile { return InputFileNoUpload(fileID) }

func FileFromURL(url string) InputFile { return InputFileNoUpload(url) }

type InputFileWithUpload struct {
	Name   string
	Reader io.Reader
}

func (InputFileWithUpload) NeedsUpload() bool { return true }

func FileFromReader(name string, reader io.Reader) InputFile {
	return &InputFileWithUpload{Name: name, Reader: reader}
}
