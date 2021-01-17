package property

import (
	"strings"
)

const (
	// Extension
	// TODO: windows
	FILE = iota // file://, /, ./ , //
	HTTP        // http://
	OS          // os://.env
	FTP         // ftp://
	// Protocol
	STRING = iota + 100 // .env, .properties
	JSON                // .json
	YAML                // .yaml, .yml
	XML                 // .xml
)

// GetProtocolAndExtension get protocol and extension
func GetProtocolAndExtension(uri string) (protocol, extension int) {
	protocol = FILE
	extension = STRING

	if strings.HasPrefix(uri, "os") {
		protocol = OS
	}
	if strings.HasPrefix(uri, "http") {
		protocol = HTTP
	}
	if strings.HasPrefix(uri, "ftp") {
		protocol = FTP
	}

	if strings.HasSuffix(uri, "json") {
		extension = JSON
	}

	if strings.HasSuffix(uri, "yaml") || strings.HasSuffix(uri, "yml") {
		extension = YAML
	}

	if strings.HasSuffix(uri, "xml") {
		extension = XML
	}
	return
}
