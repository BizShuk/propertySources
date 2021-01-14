package impl

import (
	"os"
	pSource "propertysources/propertysource"
)

// PropertiesFromOsEnv load environment variables
type PropertiesFromOsEnv map[string]string

func (p PropertiesFromOsEnv) load() {
	for _, envLine := range os.Environ() {
		key, val := parseProperty(envLine)
		p.Set(key, val)
	}
}

// Get property with key
func (p PropertiesFromOsEnv) Get(key string) (val string, ok bool) {
	val, ok = p[key]
	return
}

// Set property with key, value
func (p PropertiesFromOsEnv) Set(key string, val string) {
	p[key] = val
}

// NewPropertiesFromOsEnv PropertiesFromOsEnv constructor
func NewPropertiesFromOsEnv() pSource.Properties {
	poe := PropertiesFromOsEnv(make(map[string]string))
	poe.load()
	var p pSource.Properties = poe
	return p
}
