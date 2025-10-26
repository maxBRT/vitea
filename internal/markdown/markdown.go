package markdown

import (
	"bytes"
	"io"

	"github.com/yuin/goldmark"
)

func ExtractHTML(file io.Reader) ([]byte, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	md := goldmark.New()

	if err := md.Convert(content, buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
