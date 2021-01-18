package property

import (
	"bufio"
	"io"
	"log"
	propertyio "propertychain/io"
	"strings"
)

// OsEnvProperties load properties from file
type OsEnvProperties map[string]string

func (p OsEnvProperties) ReadFrom(reader io.Reader) (n int64, err error) {
	bufReader := bufio.NewReader(reader)

	for {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		s := string(line)
		i := strings.IndexAny(s, "=")
		p[s[0:i]] = s[i+1:]
	}

	return 0, nil
}

// Get property with key
func (p OsEnvProperties) Get(key string) (val string, ok bool) {
	val, ok = p[key]
	return
}

// Set property with key, value
func (p OsEnvProperties) Set(key string, val string) {
	p[key] = val
}

// NewOsEnvProperties NewOsEnvProperties constructor
func NewOsEnvProperties(uri string) (Properties, error) {

	reader := propertyio.GetProtocolHandler(uri)()
	defer reader.Close()

	p := OsEnvProperties(make(map[string]string))
	_, err := p.ReadFrom(reader)
	if err != nil {
		log.Fatal("Load properties failed")
		return nil, err
	}

	return p, err
}
