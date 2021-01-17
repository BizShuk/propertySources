package main

import (
	"log"
	. "propertychain/property"
)

// PropertiesChain contain multiple properties from default chaining
type PropertiesChain struct {
	prop []Properties
}

func (pc *PropertiesChain) load() {
	pc.appendProperties(CreateProperties("os://.env"))
	pc.appendProperties(CreateProperties(".env"))
}

func (pc *PropertiesChain) appendProperties(p Properties) {
	if pc == nil {
		return
	}
	pc.prop = append(pc.prop, p)
}

// Get get value with key from propertysources chain
func (pc *PropertiesChain) Get(key string) (val string, ok bool) {
	for _, v := range pc.prop {
		sVal, ok := v.Get(key)
		if ok {
			val = sVal
			return val, true
		}
	}
	return val, false
}

// New PropertySources constructor
func New() *PropertiesChain {
	pc := &PropertiesChain{}
	pc.load()
	return pc
}

func main() {
	pc := New()
	log.Println(pc.Get("a"))
}
