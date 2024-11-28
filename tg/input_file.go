package tg

import "io"

// InputFile represents a file that can be used as input.
type InputFile struct {
	Value  string    // The value of the file (e.g., file ID, URL, or path)
	Reader io.Reader // The reader for the file content
}

// IsInputFile is a marker method to indicate that the struct is an InputFile.
func (InputFile) IsInputFile() {}

// IsUploadable checks if the InputFile is uploadable.
// An InputFile is considered uploadable if it has a non-nil Reader.
func (ifu *InputFile) IsUploadable() bool {
	return ifu.Reader != nil
}

// MarshalJSON converts the InputFile to JSON.
// If the InputFile has a non-nil Reader, it returns a JSON string with an attachment reference.
// Otherwise, it returns a JSON string with the file value.
func (ifu *InputFile) MarshalJSON() ([]byte, error) {
	if ifu.Reader != nil {
		return []byte(`"attach://` + ifu.Value + `"`), nil
	}
	return []byte(`"` + ifu.Value + `"`), nil
}
