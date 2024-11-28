package tgo

import (
	"io"

	"github.com/haashemi/tgo/tg"
)

// FileFromID creates an InputFile from a file ID.
func FileFromID(fileID string) *tg.InputFile {
	return &tg.InputFile{Value: fileID}
}

// FileFromURL creates an InputFile from a URL.
func FileFromURL(url string) *tg.InputFile {
	return &tg.InputFile{Value: url}
}

// FileFromReader creates an InputFile from a name and a reader.
func FileFromReader(name string, reader io.Reader) *tg.InputFile {
	return &tg.InputFile{Value: name, Reader: reader}
}

// FileFromPath creates an InputFile from a file path.
//
// Note that you should not use this if you plan to use the public telegram bot API.
// This is only available for locally hosted bot API servers.
func FileFromPath(path string) *tg.InputFile {
	return &tg.InputFile{Value: "file://" + path}
}
