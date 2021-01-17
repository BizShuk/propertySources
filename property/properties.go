package property

import (
	"io"
	. "propertychain/utils"
)

// Properties interface for properties map
type Properties interface {
	Get(key string) (val string, ok bool)
	Set(key string, val string)
}

func CreateProperties(uri string) Properties {
	protocol, extension := GetProtocolAndExtension(uri)

	loadSourceHandler := GetPropertySourceHandler(protocol)
	createPropertyHandler := GetPropertyHandler(extension)

	sourceReader := loadSourceHandler(uri)
	p, err := createPropertyHandler(sourceReader)
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

type propertiesHandler func(reader io.Reader) (Properties, error)

var OsEnvPropertiesCreator = func(reader io.Reader) (Properties, error) {
	p, err := NewOsEnvProperties(reader)
	return p, err
}

var StringPropertiesCreator = func(reader io.Reader) (Properties, error) {
	p, err := NewStringProperties(reader)
	return p, err
}
