package impl

import (
	"bufio"
	"os"
	pSource "propertysources/propertysource"
)

// PropertiesFromFile load properties from file
type PropertiesFromFile map[string]string

func (p PropertiesFromFile) load(filename string) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	file, err := os.Open(path + "/" + filename)
	defer file.Close()
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		key, val := parseProperty(string(line))
		p.Set(key, val)
	}
	return nil
}

// Get property with key
func (p PropertiesFromFile) Get(key string) (val string, ok bool) {
	val, ok = p[key]
	return
}

// Set property with key, value
func (p PropertiesFromFile) Set(key string, val string) {
	p[key] = val
}

// NewPropertiesFromFile NewPropertiesFromFile constructor
func NewPropertiesFromFile(filename string) pSource.Properties {
	pf := PropertiesFromFile(make(map[string]string))

	err := pf.load(filename)
	if err != nil {
		return nil
	}

	var p pSource.Properties = pf
	return p
}
