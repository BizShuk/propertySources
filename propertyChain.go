package main

import (
	"log"
	. "propertychain/property"
)

// PropertiesChain contain multiple properties from default chaining
type PropertiesChain struct {
	plist []Properties
}

func (p *PropertiesChain) load() {
	p.appendProperties(CreateProperties("os://.env"))
	p.appendProperties(CreateProperties(".env"))
}

func (p *PropertiesChain) appendProperties(prop Properties) {
	if p == nil {
		return
	}
	p.plist = append(p.plist, prop)
}

// Get get value with key from propertysources chain
func (p *PropertiesChain) Get(key string) (val string, ok bool) {
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
func New() *PropertiesChain {
	p := &PropertiesChain{}
	p.load()
	return p
}

func main() {
	p := New()
	log.Println(p.Get("a"))
}
