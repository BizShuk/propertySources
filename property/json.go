package property

import (
	"bufio"
	"io"
	"log"
	propertyio "propertychain/io"
	"strings"
)

// JsonProperties load properties from file
type JsonProperties map[string]string

func (p JsonProperties) load(reader io.ReadCloser) error {
	bufReader := bufio.NewReader(reader)

	for {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		s := string(line)
		i := strings.IndexAny(s, "=")
		p[s[0:i]] = s[i+1:]
	}

	return nil
}

// Get property with key
func (p JsonProperties) Get(key string) (val string, ok bool) {
	val, ok = p[key]
	return
}

// Set property with key, value
func (p JsonProperties) Set(key string, val string) {
	p[key] = val
}

// NewJsonProperties NewJsonProperties constructor
func NewJsonProperties(uri string) (Properties, error) {

	reader := propertyio.GetProtocolHandler(uri)()
	defer reader.Close()

	pf := JsonProperties(make(map[string]string))

	err := pf.load(reader)
	if err != nil {
		log.Fatal("Load JsonProperties failed")
		return nil, err
	}
	return pf, nil
}
