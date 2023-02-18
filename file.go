package tgo

import (
	"io"

	"github.com/haashemi/tgo/botapi"
)

func FileFromID(fileID string) botapi.InputFile {
	return botapi.FileFromID(fileID)
}

func FileFromURL(url string) botapi.InputFile {
	return botapi.FileFromURL(url)
}

func FileFromReader(name string, reader io.Reader) botapi.InputFile {
	return botapi.FileFromReader(name, reader)
}
