package tgo

import (
	"io"
)

type ParseMode string

const (
	ParseModeNone       ParseMode = ""
	ParseModeMarkdown   ParseMode = "Markdown"
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
	ParseModeHTML       ParseMode = "HTML"
)

type Username string

func (Username) IsChatID() {}

type ID int64

func (ID) IsChatID() {}

type InputFile struct {
	Value  string
	Reader io.Reader
}

func (InputFile) IsInputFile() {}

func (ifu *InputFile) IsUploadable() bool {
	return ifu.Reader != nil
}

func (ifu *InputFile) MarshalJSON() ([]byte, error) {
	return []byte(`"attach://` + ifu.Value + `"`), nil
}

func FileFromID(fileID string) *InputFile {
	return &InputFile{Value: fileID}
}

func FileFromURL(url string) *InputFile {
	return &InputFile{Value: url}
}

func FileFromReader(reader io.Reader) *InputFile {
	return &InputFile{Reader: reader}
}
