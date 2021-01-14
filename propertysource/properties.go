package propertysource

// Properties interface for properties map
type Properties interface {
	Get(key string) (val string, ok bool)
	Set(key string, val string)
}
