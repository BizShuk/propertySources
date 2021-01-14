package impl

import "strings"

func parseProperty(s string) (key string, val string) {
	i := strings.IndexAny(s, "=")
	return s[0:i], s[i+1:]
}
