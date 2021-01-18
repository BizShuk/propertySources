package ui

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	. "propertychain/utils"
	"strings"
)

// Handler handle source
type Handler func() (readCloser io.ReadCloser)

// GetPropertySourceHandler factory
func GetProtocolHandler(uri string) (handler Handler) {
	protocol := GetProtocol(uri)
	switch protocol {
	case OS:
		handler = loadFromOS(uri)
	case HTTP:
		handler = loadFromHttp(uri)
	case FTP:
		handler = loadFromFtp(uri)
	default:
		handler = loadFromFile(uri)
	}
	return
}

var loadFromOS = func(uri string) Handler {
	return func() (readCloser io.ReadCloser) {
		lines := os.Environ()
		reader := strings.NewReader(strings.Join(lines, "\n"))
		readCloser = ioutil.NopCloser(reader)
		return
	}
}

var loadFromFile = func(uri string) Handler {
	return func() (readCloser io.ReadCloser) {
		readCloser, err := os.Open(uri)
		if err != nil {
			log.Fatal("File not found!!")
			return nil
		}

		return readCloser
	}
}

var loadFromHttp = func(uri string) Handler {
	return func() (readCloser io.ReadCloser) {
		resp, err := http.Get(uri)
		if err != nil {
			log.Fatal("Erro when fetch ", uri)
			return nil
		}
		readCloser = resp.Body
		return
	}
}

var loadFromFtp = func(uri string) Handler {
	return func() (readCloser io.ReadCloser) {
		readCloser, err := os.Open(uri)
		if err != nil {
			log.Fatal("File not found!!")
			return nil
		}

		return readCloser
	}
}
