package time

import (
	"errors"
)

func tzData(name string) ([]byte, bool) {
	data, ok := files["zoneinfo/"+name]
	return data, ok
}

// Address https://github.com/golang/go/issues/21881.
func loadLocationEmbeddedFile(name string) (*Location, error) {
	if name == "" || name == "UTC" || name == "Local" {
		return LoadLocation(name)
	}
	if tzdata, ok := tzData(name); ok {
		return LoadLocationFromTZData(name, tzdata)
	}
	return nil, errors.New("unknown location " + name)
}
