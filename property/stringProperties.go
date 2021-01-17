package property

import (
	"bufio"
	"io"
	"log"
	"strings"
)

// StringProperties load properties from file
type StringProperties map[string]string

func (p StringProperties) load(reader io.Reader) error {
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
func (p StringProperties) Get(key string) (val string, ok bool) {
	val, ok = p[key]
	return
}

// Set property with key, value
func (p StringProperties) Set(key string, val string) {
	p[key] = val
}

// NewStringProperties NewStringProperties constructor
func NewStringProperties(reader io.Reader) (Properties, error) {

	pf := StringProperties(make(map[string]string))

	err := pf.load(reader)
	if err != nil {
		log.Fatal("Load StringProperties failed")
		return nil, err
	}
	return pf, nil
}
