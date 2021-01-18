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
	p, err := GetPropertyHandler(uri)()
	if err != nil {
		return nil
	}
	return p
}

// GetPropertyHandler factory
func GetPropertyHandler(uri string) (ph propertiesHandler) {
	extension := GetExtension(uri)
	switch extension {
	case JSON:
	case YAML:
	case XML:
	default: // OS, STRING as well
		ph = StringPropertiesCreator(uri)
	}
	return
}

type propertiesHandler func() (Properties, error)

var StringPropertiesCreator = func(uri string) propertiesHandler {
	return func() (Properties, error) {
		p, err := NewStringProperties(uri)
		return p, err
	}
}

var JsonPropertiesCreator = func(uri string) propertiesHandler {
	return func() (Properties, error) {
		p, err := NewJsonProperties(uri)
		return p, err
	}
}
