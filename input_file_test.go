package tgo

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Interface check
var (
	RawInputFileUploadable    InputFile = &InputFileUploadable{}
	RawInputFileNotUploadable InputFile = InputFileNotUploadable("")
)

func TestInputFileUploadable(t *testing.T) {
	got := RawInputFileUploadable.NeedsUpload()

	assert.True(t, got)
}

func TestInputFileNotUploadable(t *testing.T) {
	got := RawInputFileNotUploadable.NeedsUpload()

	assert.False(t, got)
}

func TestFileFromID(t *testing.T) {
	expected := InputFileNotUploadable("some-id-123")
	got := FileFromID("some-id-123")

	assert.Equal(t, expected, got)
}

func TestFileFromURL(t *testing.T) {
	expected := InputFileNotUploadable("https://example.com/image.png")
	got := FileFromURL("https://example.com/image.png")

	assert.Equal(t, expected, got)
}

func TestFileFromReader(t *testing.T) {
	reader := bytes.NewReader(nil)

	expected := &InputFileUploadable{
		Name:   "SomeName",
		Reader: reader,
	}
	got := FileFromReader("SomeName", reader)

	assert.Equal(t, expected, got)
}
