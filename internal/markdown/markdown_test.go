package markdown

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractHTML(t *testing.T) {
	// OK
	t.Run("Extract HTML from Markdown file", func(t *testing.T) {
		file := []byte("# Hello World")
		buffer, err := ExtractHTML(bytes.NewReader(file))
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "<h1>Hello World</h1>\n", string(buffer))
	})
	// Empty file
	t.Run("Extract HTML from invalid Markdown file", func(t *testing.T) {
		var file []byte
		_, err := ExtractHTML(bytes.NewReader(file))
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "", string(file))
	})
}
