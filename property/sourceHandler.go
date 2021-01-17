package property

import (
	"io"
	"log"
	"os"
	. "propertychain/utils"
	"strings"
)

type sourceHandler func(uri string) (reader io.Reader)

// GetPropertySourceHandler factory
func GetPropertySourceHandler(protocol int) (loadSourceHandler sourceHandler) {
	switch protocol {
	case OS:
		loadSourceHandler = loadFromOS
	case FILE:
		loadSourceHandler = loadFromFile
	default:
	}

	return
}

var loadFromOS = func(uri string) (reader io.Reader) {
	lines := os.Environ()
	reader = strings.NewReader(strings.Join(lines, "\n"))
	return
}

var loadFromFile = func(uri string) (reader io.Reader) {
	reader, err := os.Open(uri)
	if err != nil {
		log.Fatal("File not found!!")
		return nil
	}
	return
}
