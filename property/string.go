package property

import (
	"bufio"
	"io"
	"log"
	propertyio "propertychain/io"
	"strings"
)

// StringProperties load properties from file
type StringProperties map[string]string

func (p StringProperties) ReadFrom(reader io.Reader) (n int64, err error) {
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
func (p StringProperties) Get(key string) (val string, ok bool) {
	val, ok = p[key]
	return
}

// Set property with key, value
func (p StringProperties) Set(key string, val string) {
	p[key] = val
}

// NewStringProperties NewStringProperties constructor
func NewStringProperties(uri string) (Properties, error) {

	reader := propertyio.GetProtocolHandler(uri)()
	defer reader.Close()

	p := StringProperties(make(map[string]string))

	_, err := p.ReadFrom(reader)
	if err != nil {
		log.Fatal("Load StringProperties failed")
		return nil, err
	}

	return p, nil
}
