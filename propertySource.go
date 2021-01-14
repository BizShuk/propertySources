package main

import (
	pSource "propertysources/propertysource"
	"propertysources/propertysource/impl"
)

// PropertiesSources contain multiple properties from default chaining
type PropertiesSources struct {
	plist []pSource.Properties
}

func (p *PropertiesSources) load() {
	p.appendProperties(impl.NewPropertiesFromOsEnv())
	p.appendProperties(impl.NewPropertiesFromFile(".env"))
}

func (p *PropertiesSources) appendProperties(properties pSource.Properties) {
	if p == nil {
		return
	}
	p.plist = append(p.plist, properties)
}

// Get get value with key from propertysources chain
func (p *PropertiesSources) Get(key string) (val string, ok bool) {
	for _, v := range p.plist {
		sVal, ok := v.Get(key)
		if ok {
			val = sVal
			return val, true
		}
	}
	return val, false
}

// New PropertySources constructor
func New() *PropertiesSources {
	p := &PropertiesSources{}
	p.load()
	return p
}
