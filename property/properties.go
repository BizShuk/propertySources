package property

import (
	"io"
	. "propertychain/utils"
)

// Properties interface for properties map
type Properties interface {
	Get(key string) (val string, ok bool)
	Set(key string, val string)
	io.ReaderFrom
}

func CreateProperties(uri string) Properties {
	extension := GetExtension(uri)
	createPropertyHandler := GetPropertyHandler(extension)

	p, err := createPropertyHandler(uri)
	if err != nil {
		return nil
	}
	return p
}

// GetPropertyHandler factory
func GetPropertyHandler(extension int) (ph propertiesHandler) {
	switch extension {
	case OS:
		ph = OsEnvPropertiesCreator
	case STRING:
		ph = StringPropertiesCreator
	case JSON:
	case YAML:
	case XML:
	default:
	}
	return
}

type propertiesHandler func(uri string) (Properties, error)

var OsEnvPropertiesCreator = func(uri string) (Properties, error) {
	p, err := NewOsEnvProperties(uri)
	return p, err
}

var StringPropertiesCreator = func(uri string) (Properties, error) {
	p, err := NewStringProperties(uri)
	return p, err
}

var JsonPropertiesCreator = func(uri string) (Properties, error) {
	p, err := NewJsonProperties(uri)
	return p, err
}
