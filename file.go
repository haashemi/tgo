package tgo

import "io"

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

func FileFromReader(name string, reader io.Reader) *InputFile {
	return &InputFile{Value: name, Reader: reader}
}

func FileFromPath(path string) *InputFile {
	return &InputFile{Value: "file://" + path}
}
